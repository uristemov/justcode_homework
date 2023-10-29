package usecase

import (
	"context"
	"homeworks/hw12/internal/auth/auth/repository"
	"homeworks/hw12/internal/auth/config"
	"homeworks/hw12/internal/auth/entity"
	"homeworks/hw12/internal/auth/pkg/jwttoken"
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
	CreateUser(ctx context.Context, u *entity.User) (string, error)
	Login(ctx context.Context, email, password string) (string, error)
	GetUser(ctx context.Context, id string) (*entity.User, error)
}
