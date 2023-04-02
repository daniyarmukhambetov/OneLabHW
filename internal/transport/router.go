package transport

import (
	"github.com/go-chi/chi/v5"
	"hw1/internal/config"
	"hw1/internal/handler"
)

func NewRouter(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()
	h := handler.NewHandler(cfg)
	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.User.List)
		r.Get("/{id}", h.User.Retrieve)
		r.Post("/", h.User.Create)
	})
	return r
}
