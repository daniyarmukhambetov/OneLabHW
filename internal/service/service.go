package service

import (
	"fmt"
	"hw1/config"
	"hw1/internal/models"
	"hw1/internal/storage"
)

type Manager struct {
	User     IUserService
	Book     IBookService
	BookUser IBookUserService
}

func NewManager(storage *storage.Storage, cfg *config.Config) (*Manager, error) {
	if storage == nil {
		return nil, fmt.Errorf("storage is empty")
	}
	return &Manager{User: NewUserService(storage, cfg), Book: NewBook(storage), BookUser: NewBookUser(storage)}, nil
}

type IUserService interface {
	List() ([]models.User, error)
	Retrieve(string) (models.User, error)
	Create(models.UserModelIn) (models.User, error)
	Update(string, models.UserUpdate) (models.User, error)
	Delete(string) (string, error)
	GetJWT(string, string) (models.JWT, error)
}

type IBookService interface {
	List() ([]models.Book, error)
	Retrieve(string, error) (models.Book, error)
	Create(models.Book, error) (models.Book, error)
	Update(models.Book, error) (models.Book, error)
	Delete(string, error) (string, error)
}

type IBookUserService interface {
	List() ([]models.BookUser, error)
	ListUserBookCount() ([]models.UserBookCount, error)
}
