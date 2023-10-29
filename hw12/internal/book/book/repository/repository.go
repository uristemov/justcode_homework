package repository

import (
	"context"
	"homeworks/hw12/internal/book/entity"
)

type Repository interface {
	GetAllBooks(ctx context.Context) ([]entity.Book, error)
	GetBookById(ctx context.Context, id string) (*entity.Book, error)
}
