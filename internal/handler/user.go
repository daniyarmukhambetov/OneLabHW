package handler

import (
	"github.com/go-chi/chi/v5"
	"hw1/internal/config"
	"hw1/internal/logger"
	"hw1/internal/models"
	"hw1/internal/service"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Service *service.Manager
}

func (h *UserHandler) List(writer http.ResponseWriter, request *http.Request) {
	users := h.Service.User.List()
	err := writeJSON(writer, http.StatusOK, users, nil)
	if err != nil {
		logger.Logger().Println("server error", err)
		serverErrorResponse(writer, request, err)
	}
}

func (h *UserHandler) Retrieve(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		logger.Logger().Println("error", err)
		errorResponse(writer, request, http.StatusBadRequest, "cannot parse id") // тут молодец, предусмотрел ошибки
		return
	}
	var user models.UserModel
	user, err = h.Service.User.Retrieve(id)
	if err != nil {
		errorResponse(writer, request, http.StatusBadRequest, err)
		return
	}
	err = writeJSON(writer, http.StatusOK, user, nil)
	if err != nil {
		logger.Logger().Println("server error", err)
		serverErrorResponse(writer, request, err)
	}
}

func (h *UserHandler) Create(writer http.ResponseWriter, request *http.Request) {
	var userIn models.UserModelIn
	err := readJSON(writer, request, &userIn)
	if err != nil {
		errorResponse(writer, request, http.StatusBadRequest, err)
		return
	}
	var user models.UserModel
	user, err = h.Service.User.Create(userIn)
	if err != nil {
		errorResponse(writer, request, http.StatusBadRequest, err)
		return
	}
	err = writeJSON(writer, http.StatusCreated, user, nil)
	if err != nil {
		logger.Logger().Println("server error", err)
		serverErrorResponse(writer, request, err)
	}
}

func NewUserHandler(cfg *config.Config, s *service.Manager) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}
