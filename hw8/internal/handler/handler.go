package handler

import "homeworks/hw8/internal/service"

type Handler struct {
	service service.Service
}

func New(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
