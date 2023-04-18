package service

import (
	"fmt"
	"hw1/config"
	"hw1/logger"
	"hw1/models"
	"hw1/pkg"
	"hw1/storage"
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

func (s *UserService) List() ([]models.User, error) {
	return s.Repo.User.List()
}

func (s *UserService) Retrieve(username string) (models.User, error) {
	res, err := s.Repo.User.Retrieve(username)
	if err != nil {
		return models.User{}, err
	}
	return res, nil
}

func (s *UserService) Create(in models.UserModelIn) (models.User, error) {
	hash, err := pkg.HashPassword(in.Password)
	if err != nil {
		logger.Logger().Println(err)
		return models.User{}, err
	}
	in.Password = hash
	return s.Repo.User.Create(in)
}

func (s *UserService) Update(i string, in models.UserUpdate) (models.User, error) {
	hash, err := pkg.HashPassword(in.Password)
	if err != nil {
		logger.Logger().Println(err)
		return models.User{}, err
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

type IUserService interface {
	List() ([]models.User, error)
	Retrieve(string) (models.User, error)
	Create(models.UserModelIn) (models.User, error)
	Update(string, models.UserUpdate) (models.User, error)
	Delete(string) (string, error)
	GetJWT(string, string) (models.JWT, error)
}
