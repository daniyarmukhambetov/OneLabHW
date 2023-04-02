package storage

import (
	"hw1/internal/models"
	"hw1/internal/storage/in_memory"
)

type Storage struct {
	User IUser
}

func NewStorage() *Storage {
	return &Storage{User: in_memory.NewUserRepo()}
}

type IUser interface {
	List() []models.UserModel
	Retrieve(int) (models.UserModel, error)
	Create(models.UserModelIn) (models.UserModel, error)
	Update(int, models.UserModelIn) (models.UserModel, error)
	Delete(int) (int, error)
}
