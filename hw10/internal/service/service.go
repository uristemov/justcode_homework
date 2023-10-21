package service

import (
	"context"
	"homeworks/hw10/internal/entity"
)

type Service interface {
	GetUser(ctx context.Context, id string) (*entity.User, error)
}
