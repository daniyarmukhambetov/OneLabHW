package handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"hw1/logger"
	"hw1/models"
	"hw1/service"
	"log"
	"net/http"
)

type UserHandler struct {
	Service *service.Manager
}

// @Summary      List users
// @Description  get users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.User
// @Router       /users [get]
// @Failure 	 500
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

// @Summary      Update user
// @Description  Update User
// @Tags         users
// @Accept       json
// @Produce      json
// @Param 		 username path string true "username"
// @Success      200  {object}   models.User
// @Failure      400
// @Router       /users/{username} [get]
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

// @Summary      create user
// @Description  create User
// @Tags         users
// @Accept       json
// @Produce      json
// @Param 		 data body models.UserModelIn true "create user"
// @Success      200  {object}   models.User
// @Failure      400
// @Router       /users [post]
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

// @Summary      Update user
// @Description  Update User
// @Tags         users
// @Accept       json
// @Produce      json
// @Param 		 data body models.UserUpdate true "update users name, last_name and password"
// @Param 		 username path string true "username"
// @Success      200  {object}   models.User
// @Failure      400
// @Router       /users/{username} [put]
func (h *UserHandler) Update(writer http.ResponseWriter, request *http.Request) {
	username := request.Context().Value("user").(string)
	fmt.Println(username)
	//fmt.Println(request.Context())
	//username := "usr"
	var usr models.User
	usr, err := h.Service.User.Retrieve(username)
	if err != nil {
		errorResponse(writer, request, http.StatusBadRequest, err.Error())
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

// @Summary      Get JWT
// @Description  Log in
// @Tags         users
// @Accept       json
// @Produce      json
// @Param 		 data body models.Login true "login"
// @Success      200  {object}   models.JWT
// @Failure      400
// @Router       /users/login [post]
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

func NewUserHandler(s *service.Manager) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}
