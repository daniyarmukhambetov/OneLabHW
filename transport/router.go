package transport

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"hw1/handler"
)

func NewRouter(h *handler.Handler) *chi.Mux {
	r := chi.NewRouter()
	//r.Use(MyMiddleware)
	r.Route("/api", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(MyMiddleware)
			r.Put("/users/{username}", h.User.Update)
			r.Post("/rents", h.Rent.RentBook)
		})
		r.Get("/users/{username}", h.User.Retrieve)
		r.Get("/users", h.User.List)
		r.Post("/users", h.User.Create)
		r.Post("/users/login", h.User.GetJWT)
		r.Get("/rents", h.Rent.ListUserRentedBooks)
	})
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://127.0.0.1:8080/swagger/doc.json"), //The url pointing to API definition
	))
	return r
}
