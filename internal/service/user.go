package service

import (
	"hw1/internal/models"
	"hw1/internal/storage"
)

type UserService struct {
	Repo *storage.Storage
}

func NewUserService(cfg int, repo *storage.Storage) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) List() []models.UserModel {
	return s.Repo.User.List()
}

func (s *UserService) Retrieve(i int) (models.UserModel, error) {
	res, err := s.Repo.User.Retrieve(i)
	if err != nil {
		return models.UserModel{}, err
	}
	// do something, mb validating and errors

	return res, nil
}

func (s *UserService) Create(in models.UserModelIn) (models.UserModel, error) {
	return s.Repo.User.Create(in)
}

func (s *UserService) Update(i int, in models.UserModelIn) (models.UserModel, error) {
	return s.Repo.User.Update(i, in)
}

func (s *UserService) Delete(i int) (int, error) {
	return s.Repo.User.Delete(i)
}
