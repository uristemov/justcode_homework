package handler

import (
	"homeworks/hw10/internal/service"
	"homeworks/hw10/pkg/cache"
)

type Handler struct {
	srvs      service.Service
	userCache cache.UserCache
}

func New(srvs service.Service, userCache cache.UserCache) *Handler {
	return &Handler{
		srvs:      srvs,
		userCache: userCache,
	}
}
