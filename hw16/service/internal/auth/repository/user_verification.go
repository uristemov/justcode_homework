package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func (r *Repo) CreateUserCode(ctx context.Context, usercode string, login string) error {
	q := `
INSERT INTO user_codes (user_code, login)
VALUES (?, ?);`
	query, args, err := sqlx.In(
		q,
		usercode,
		login,
	)

	if err != nil {
		return fmt.Errorf("query bake failed: %w", err)
	}

	_, err = r.main.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("db exec query failed: %w", err)
	}

	return nil
}

func (r *Repo) GetUserCode(ctx context.Context, login string) (string, error) {
	q := ` SELECT user_code FROM user_codes WHERE login=?;`
	query, args, err := sqlx.In(
		q,
		login,
	)

	if err != nil {
		return "", fmt.Errorf("query bake failed: %w", err)
	}

	var userCode string
	result, err := r.main.QueryxContext(ctx, query, args...)
	result.Scan(&userCode)

	if err != nil {
		return "", fmt.Errorf("db exec query failed: %w", err)
	}

	return userCode, nil
}
