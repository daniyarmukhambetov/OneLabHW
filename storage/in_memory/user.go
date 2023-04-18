package in_memory

import (
	"hw1/logger"
	"hw1/models"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) List() []models.User {
	return data
}

func (r *UserRepo) Retrieve(i int) (models.User, error) {
	for _, user := range data {
		if user.ID == i {
			return user, nil
		}
	}
	logger.Logger().Println("db error", IDNotFoundError)
	return models.User{}, IDNotFoundError
}

func (r *UserRepo) Create(u models.UserModelIn) (models.User, error) {
	for _, user := range data {
		if user.Email == u.Email {
			logger.Logger().Println("db error", EmailDuplicationError)
			return models.User{}, EmailDuplicationError
		}
		if user.Username == u.Username {
			logger.Logger().Println("db error", UsernameDuplicationError)
			return models.User{}, UsernameDuplicationError
		}
	}
	data = append(data,
		models.User{
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

func (r *UserRepo) Update(i int, u models.UserModelIn) (models.User, error) {
	for _, user := range data {
		if user.Email == u.Email {
			logger.Logger().Println("db error", EmailDuplicationError)
			return models.User{}, EmailDuplicationError
		}
		if user.Username == u.Username {
			logger.Logger().Println("db error", UsernameDuplicationError)
			return models.User{}, UsernameDuplicationError
		}
	}
	for _, user := range data {
		if user.ID == i {
			user = models.User{
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
	return models.User{}, IDNotFoundError
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
