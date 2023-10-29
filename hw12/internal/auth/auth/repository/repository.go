package repository

import (
	"context"
	"homeworks/hw12/internal/auth/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, u *entity.User) (string, error)
	GetUser(ctx context.Context, email string) (*entity.User, error)
}
