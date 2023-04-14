package service

import (
	"fmt"
	"hw1/config"
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
