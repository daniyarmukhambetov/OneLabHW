package handler

import (
	"hw1/internal/config"
	service2 "hw1/internal/service"
	"net/http"
)

type Handler struct {
	User IUserHandler
}

func NewHandler(cfg *config.Config) *Handler {
	service := service2.NewManager(cfg.DBCfg.DbType)
	return &Handler{
		User: NewUserHandler(cfg, service),
	}
}

type IUserHandler interface {
	List(http.ResponseWriter, *http.Request)
	Retrieve(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
}
