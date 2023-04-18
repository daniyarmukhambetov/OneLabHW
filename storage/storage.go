package storage

import (
	"gorm.io/gorm"
	"hw1/models"
	"hw1/storage/postgres"
)

type Storage struct {
	User     IUser
	Book     IBook
	BookUser IBookUser
}

func NewStorage(db *gorm.DB) (*Storage, error) {
	return &Storage{
		User:     postgres.NewUser(db),
		Book:     postgres.NewBook(db),
		BookUser: postgres.NewBookUser(db),
	}, nil
}

type IUser interface {
	List() ([]models.User, error)
	Retrieve(string) (models.User, error)
	Create(models.UserModelIn) (models.User, error)
	Update(string, models.UserUpdate) (models.User, error)
	Delete(string) (string, error)
}

type IBook interface {
	List() ([]models.Book, error)
	Retrieve(name string) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Update(book models.Book) (models.Book, error)
	Delete(name string) (string, error)
}

type IBookUser interface {
	List() ([]models.BookUser, error)
	ListUserBookCount(string) ([]models.UserBookCount, error)
	Create(string, int) (models.BookUser, error)
}
