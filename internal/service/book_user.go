package service

import (
	"hw1/internal/models"
	"hw1/internal/storage"
)

type BookUser struct {
	Repo *storage.Storage
}

func (s *BookUser) List() ([]models.BookUser, error) {
	return s.Repo.BookUser.List()
}

func (s *BookUser) ListUserBookCount() ([]models.UserBookCount, error) {
	date := "2002.02.02"
	res, err := s.Repo.BookUser.ListUserBookCount(date)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func NewBookUser(repo *storage.Storage) *BookUser {
	return &BookUser{Repo: repo}
}

type IBookUserService interface {
	List() ([]models.BookUser, error)
	ListUserBookCount() ([]models.UserBookCount, error)
}
