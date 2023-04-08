package service

import (
	"fmt"
	"hw1/internal/config"
	"hw1/internal/logger"
	"hw1/internal/models"
	"hw1/internal/storage"
	"hw1/pkg"
	"log"
)

type UserService struct {
	Repo *storage.Storage
	Cfg  *config.Config
}

func NewUserService(repo *storage.Storage, cfg *config.Config) *UserService {
	return &UserService{
		Repo: repo,
		Cfg:  cfg,
	}
}

func (s *UserService) List() ([]models.UserModel, error) {
	return s.Repo.User.List()
}

func (s *UserService) Retrieve(username string) (models.UserModel, error) {
	res, err := s.Repo.User.Retrieve(username)
	if err != nil {
		return models.UserModel{}, err
	}
	return res, nil
}

func (s *UserService) Create(in models.UserModelIn) (models.UserModel, error) {
	hash, err := pkg.HashPassword(in.Password)
	if err != nil {
		logger.Logger().Println(err)
		return models.UserModel{}, err
	}
	in.Password = hash
	return s.Repo.User.Create(in)
}

func (s *UserService) Update(i string, in models.UserUpdate) (models.UserModel, error) {
	hash, err := pkg.HashPassword(in.Password)
	if err != nil {
		logger.Logger().Println(err)
		return models.UserModel{}, err
	}
	in.Password = hash
	return s.Repo.User.Update(i, in)
}

func (s *UserService) GetJWT(username string, password string) (models.JWT, error) {
	usr, err := s.Repo.User.Retrieve(username)
	if err != nil {
		log.Println(err)
		return models.JWT{}, err
	}
	var JWT models.JWT
	if usr.ID == 0 {
		log.Println("no user")
		return models.JWT{}, fmt.Errorf("incorrect username")
	}
	if !pkg.CheckPasswordHash(password, usr.Password) {
		fmt.Println("password mismatch")
		return models.JWT{}, fmt.Errorf("incorrect password")
	}
	token, err := pkg.GenerateJWT(username, s.Cfg.JWTSecret)
	if err != nil {
		log.Println(err)
		return models.JWT{}, err
	}
	fmt.Sprintf(token)
	JWT.Token = token
	return JWT, nil
}

func (s *UserService) Delete(i string) (string, error) {
	return s.Repo.User.Delete(i)
}
