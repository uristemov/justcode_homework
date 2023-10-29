package usecase

import (
	"context"
	"homeworks/hw12/internal/book/book/repository"
	"homeworks/hw12/internal/book/config"
	"homeworks/hw12/internal/book/entity"
	"homeworks/hw12/internal/book/pkg/jwttoken"
)

type Manager struct {
	Repository repository.Repository
	Config     *config.Config
	Token      *jwttoken.JWTToken
}

func New(repository repository.Repository, config *config.Config, token *jwttoken.JWTToken) *Manager {
	return &Manager{Repository: repository, Config: config, Token: token}
}

type Service interface {
	GetAllBooks(ctx context.Context) ([]entity.Book, error)
	GetBookById(ctx context.Context, id string) (*entity.Book, error)
	VerifyToken(token string) (string, error)
}
