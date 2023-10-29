package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"homeworks/hw12/internal/book/entity"
	"log"
	"strings"
)

func (p *Postgres) GetAllBooks(ctx context.Context) ([]entity.Book, error) {
	var books []entity.Book
	query := fmt.Sprintf("SELECT id, name, genre, annotation, image_path FROM %s", bookTable)
	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		book := entity.Book{}
		err = rows.Scan(&book.Id, &book.Name, &book.Genre, &book.Annotation, &book.ImagePath)
		books = append(books, book)
		if err != nil {
			log.Printf("Scan book values error %s", err.Error())
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (p *Postgres) GetBookById(ctx context.Context, id string) (*entity.Book, error) {

	book := new(entity.Book)

	query := fmt.Sprintf("SELECT id, name,genre, annotation, image_path FROM %s WHERE id='%s'", bookTable, strings.TrimSpace(id))
	err := pgxscan.Get(ctx, p.Pool, book, query)
	if err != nil {
		return nil, err
	}

	return book, nil
}
