package handler

import "homeworks/hw12/internal/book/book/usecase"

type Handler struct {
	srvs usecase.Service
}

func New(srvc usecase.Service) *Handler {
	return &Handler{
		srvs: srvc,
	}
}
