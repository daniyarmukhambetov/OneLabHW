package service

import (
	"hw1/internal/models"
	"hw1/internal/storage"
)

type Manager struct {
	User IUserService
}

func NewManager(cfg int) *Manager {
	Repo := storage.NewStorage() // перенести в main 
	return &Manager{User: NewUserService(cfg, Repo)}
}

type IUserService interface {
	List() []models.UserModel // здесь тоже может быть ошибка)
	Retrieve(int) (models.UserModel, error)
	Create(models.UserModelIn) (models.UserModel, error)
	Update(int, models.UserModelIn) (models.UserModel, error)
	Delete(int) (int, error)
}
