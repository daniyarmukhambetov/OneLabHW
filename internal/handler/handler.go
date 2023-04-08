package handler

import (
	"hw1/internal/service"
	"net/http"
)

type Handler struct {
	User IUserHandler
}

func NewHandler(service *service.Manager) (*Handler, error) {
	return &Handler{
		User: NewUserHandler(service),
	}, nil
}

type IUserHandler interface {
	List(http.ResponseWriter, *http.Request)
	Retrieve(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	GetJWT(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	GetBooks(http.ResponseWriter, *http.Request)
	ListUserBookCount(http.ResponseWriter, *http.Request)
}
