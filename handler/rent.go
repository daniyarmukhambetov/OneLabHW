package handler

import (
	_ "hw1/dto"
	"hw1/service"
	"net/http"
)

type Rent struct {
	Service *service.Manager
}

// @Summary      renting book
// @Description  auth required
// @Tags         rents
// @Accept       json
// @Produce      json
// @Param 		 data body dto.BookName true "create rent"
// @Success      200  {object}   dto.UserRented
// @Failure      400
// @Router       /rents [post]
func (h *Rent) RentBook(writer http.ResponseWriter, request *http.Request) {
	var bookName struct {
		BookName string `json:"book_name"`
	}
	err := readJSON(writer, request, &bookName)
	if err != nil {
		errorResponse(writer, request, http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.Service.Rent.RentBook(request.Context().Value("user").(string), bookName.BookName)
	if err != nil {
		errorResponse(writer, request, http.StatusBadRequest, err.Error())
		return
	}
	err = writeJSON(writer, http.StatusCreated, res, nil)
	if err != nil {
		serverErrorResponse(writer, request, err)
		return
	}
}

// @Summary      List Rents
// @Description  get rents
// @Tags         rents
// @Accept       json
// @Produce      json
// @Success 	 200 {array} dto.BookRented
// @Router       /rents [get]
// @Failure 	 500
func (h *Rent) ListUserRentedBooks(writer http.ResponseWriter, request *http.Request) {
	res, err := h.Service.Rent.RentedBooks()
	if err != nil {
		errorResponse(writer, request, http.StatusBadRequest, err.Error())
		return
	}
	err = writeJSON(writer, http.StatusOK, res, nil)
	if err != nil {
		serverErrorResponse(writer, request, err)
		return
	}
}

func NewRent(service *service.Manager) *Rent {
	return &Rent{Service: service}
}
