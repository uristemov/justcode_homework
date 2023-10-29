package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"homeworks/hw12/internal/auth/entity"
	"log"
	"time"
)

func (p *Postgres) CreateUser(ctx context.Context, u *entity.User) (string, error) {

	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	var userId string
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                email, -- 1 
			                first_name, -- 2
			                last_name, -- 3
			                password, -- 4,
							created_at
			                )
			VALUES ($1, $2, $3, $4, $5) RETURNING id
			`, usersTable)

	err = p.Pool.QueryRow(ctx, query, u.Email, u.FirstName, u.LastName, u.Password, time.Now()).Scan(&userId)
	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}

	return userId, tx.Commit(ctx)
}

func (p *Postgres) GetUser(ctx context.Context, email string) (*entity.User, error) {
	user := new(entity.User)
	//var userID string
	query := fmt.Sprintf("SELECT id, email, first_name, last_name, password FROM %s WHERE email = '%s'", usersTable, email)

	//rows, err := p.SQLDB.Query(query, username)
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//
	//for rows.Next() {
	//	err := rows.Scan(&user.ID, &user.Username, &user.LastName, &user.LastName, &user.Password)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	//err = rows.Err()
	//if err != nil {
	//	return nil, err
	//}

	err := pgxscan.Get(ctx, p.Pool, user, query)
	if err != nil {
		log.Println("Error after pgx get")
		return nil, err
	}

	return user, nil
}
