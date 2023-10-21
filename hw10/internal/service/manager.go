package service

import (
	"homeworks/hw10/internal/repository"
)

type Manager struct {
	Repository repository.Repository
	//Config     *config.Config
	//Token      *jwttoken.JWTToken
}

func New(repository repository.Repository) *Manager {
	return &Manager{Repository: repository}
}
