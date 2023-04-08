package transport

import (
	"github.com/go-chi/chi/v5"
	"hw1/internal/handler"
)

func NewRouter(h *handler.Handler) *chi.Mux {
	r := chi.NewRouter()
	//r.Use(MyMiddleware)
	r.Route("/users", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(MyMiddleware)
			r.Put("/{username}", h.User.Update)
		})
		r.Get("/{username}", h.User.Retrieve)
		r.Get("/", h.User.List)
		r.Post("/", h.User.Create)
		r.Post("/login", h.User.GetJWT)
		r.Get("/books", h.User.GetBooks)
		r.Get("/books/count", h.User.ListUserBookCount)
	})
	return r
}
