package handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"hw1/internal/logger"
	"hw1/internal/models"
	"hw1/internal/service"
	"log"
	"net/http"
)

type UserHandler struct {
	Service *service.Manager
}

func (h *UserHandler) List(writer http.ResponseWriter, request *http.Request) {
	users, err := h.Service.User.List()
	if err != nil {
		serverErrorResponse(writer, request, err)
		return
	}
	err = writeJSON(writer, http.StatusOK, users, nil)
	if err != nil {
		logger.Logger().Println("server error", err)
		serverErrorResponse(writer, request, err)
	}
}

func (h *UserHandler) Retrieve(writer http.ResponseWriter, request *http.Request) {
	username := chi.URLParam(request, "username")
	var user models.User
	user, err := h.Service.User.Retrieve(username)
	if err != nil {
		errorResponse(writer, request, http.StatusBadRequest, err.Error())
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
		errorResponse(writer, request, http.StatusBadRequest, err.Error())
		return
	}
	var user models.User
	user, err = h.Service.User.Create(userIn)
	if err != nil {
		errorResponse(writer, request, http.StatusBadRequest, err.Error())
		return
	}
	err = writeJSON(writer, http.StatusCreated, user, nil)
	if err != nil {
		logger.Logger().Println("server error", err)
		serverErrorResponse(writer, request, err)
	}
}

func (h *UserHandler) Update(writer http.ResponseWriter, request *http.Request) {
	username := request.Context().Value("user").(string)
	fmt.Println(username)
	//fmt.Println(request.Context())
	//username := "usr"
	var usr models.User
	usr, err := h.Service.User.Retrieve(username)
	if err != nil {
		serverErrorResponse(writer, request, err)
		return
	}
	if usr.ID == 0 {
		errorResponse(writer, request, http.StatusBadRequest, "username not found")
		return
	}
	var userUpdate models.UserUpdate
	err = readJSON(writer, request, &userUpdate)
	if err != nil {
		errorResponse(writer, request, http.StatusBadRequest, err.Error())
		return
	}
	var user models.User
	user, err = h.Service.User.Update(usr.Username, userUpdate)
	if err != nil {
		errorResponse(writer, request, http.StatusBadRequest, err.Error())
		return
	}
	err = writeJSON(writer, http.StatusCreated, user, nil)
	if err != nil {
		logger.Logger().Println("server error", err)
		serverErrorResponse(writer, request, err)
	}
}

func (h *UserHandler) GetJWT(writer http.ResponseWriter, request *http.Request) {
	var LogIn models.Login
	err := readJSON(writer, request, &LogIn)
	if err != nil {
		log.Println(err)
		errorResponse(writer, request, http.StatusBadRequest, err.Error())
		return
	}
	var JWT models.JWT
	JWT, err = h.Service.User.GetJWT(LogIn.Username, LogIn.Password)
	if err != nil {
		log.Println(err)
		errorResponse(writer, request, http.StatusBadRequest, err.Error())
		return
	}
	err = writeJSON(writer, http.StatusOK, JWT, nil)
	if err != nil {
		log.Println(err)
		serverErrorResponse(writer, request, err)
		return
	}
}

func (h *UserHandler) GetBooks(writer http.ResponseWriter, request *http.Request) {
	var bookUser []models.BookUser
	bookUser, err := h.Service.BookUser.List()
	if err != nil {
		serverErrorResponse(writer, request, err)
		return
	}
	err = writeJSON(writer, http.StatusOK, bookUser, nil)
	if err != nil {
		serverErrorResponse(writer, request, err)
		return
	}
}

func (h *UserHandler) ListUserBookCount(writer http.ResponseWriter, request *http.Request) {
	var userBooksCount []models.UserBookCount
	userBooksCount, err := h.Service.BookUser.ListUserBookCount()
	if err != nil {
		serverErrorResponse(writer, request, err)
		return
	}
	err = writeJSON(writer, http.StatusOK, userBooksCount, nil)
	if err != nil {
		serverErrorResponse(writer, request, err)
		return
	}
}

func NewUserHandler(s *service.Manager) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}
