package hw15

import "homeworks/hw15/service"

type Handler struct {
	srvs service.Service
}

func New(srvc service.Service) *Handler {
	return &Handler{
		srvs: srvc,
	}
}
