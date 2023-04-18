package handler

import (
	"hw1/service"
	"net/http"
)

type Handler struct {
	User IUserHandler
	Rent IRentHandler
}

func NewHandler(service *service.Manager) (*Handler, error) {
	return &Handler{
		User: NewUserHandler(service),
		Rent: NewRent(service),
	}, nil
}

type IUserHandler interface {
	List(http.ResponseWriter, *http.Request)
	Retrieve(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	GetJWT(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
}

type IRentHandler interface {
	RentBook(http.ResponseWriter, *http.Request)
	ListUserRentedBooks(http.ResponseWriter, *http.Request)
}
