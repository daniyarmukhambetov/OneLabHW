package in_memory

import (
	"hw1/internal/logger"
	"hw1/internal/models"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) List() []models.UserModel {
	return data
}

func (r *UserRepo) Retrieve(i int) (models.UserModel, error) {
	for _, user := range data {
		if user.ID == i {
			return user, nil
		}
	}
	logger.Logger().Println("db error", IDNotFoundError)
	return models.UserModel{}, IDNotFoundError
}

func (r *UserRepo) Create(u models.UserModelIn) (models.UserModel, error) {
	for _, user := range data {
		if user.Email == u.Email {
			logger.Logger().Println("db error", EmailDuplicationError)
			return models.UserModel{}, EmailDuplicationError
		}
		if user.Username == u.Username {
			logger.Logger().Println("db error", UsernameDuplicationError)
			return models.UserModel{}, UsernameDuplicationError
		}
	}
	data = append(data,
		models.UserModel{
			ID:       len(data) + 1,
			Username: u.Username,
			Email:    u.Email,
			Name:     u.Name,
			LastName: u.LastName,
			Password: u.Password,
		},
	)

	return data[len(data)-1], nil
}

func (r *UserRepo) Update(i int, u models.UserModelIn) (models.UserModel, error) {
	for _, user := range data {
		if user.Email == u.Email {
			logger.Logger().Println("db error", EmailDuplicationError)
			return models.UserModel{}, EmailDuplicationError
		}
		if user.Username == u.Username {
			logger.Logger().Println("db error", UsernameDuplicationError)
			return models.UserModel{}, UsernameDuplicationError
		}
	}
	for _, user := range data {
		if user.ID == i {
			user = models.UserModel{
				ID:       len(data) + 1,
				Username: u.Username,
				Email:    u.Email,
				Name:     u.Name,
				LastName: u.LastName,
				Password: u.Password,
			}
		}
	}
	logger.Logger().Println("db error", IDNotFoundError)
	return models.UserModel{}, IDNotFoundError
}

func (r *UserRepo) Delete(i int) (int, error) {
	for j, user := range data {
		if user.ID == i {
			id := user.ID
			data = append(data[:j], data[j+1:]...)
			return id, nil
		}
	}
	return 0, IDNotFoundError
}
