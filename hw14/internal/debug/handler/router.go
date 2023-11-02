package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func InitRouter() http.Handler {
	r := chi.NewRouter()

	r.Mount("/debug", middleware.Profiler())

	return r
}
