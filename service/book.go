package service

import (
	"hw1/models"
	"hw1/storage"
)

type Book struct {
	Repo *storage.Storage
}

func (r *Book) List() ([]models.Book, error) {
	return r.Repo.Book.List()
}

func (r *Book) Retrieve(s string, err error) (models.Book, error) {
	return r.Repo.Book.Retrieve(s)
}

func (r *Book) Create(book models.Book, err error) (models.Book, error) {
	return r.Repo.Book.Create(book)
}

func (r *Book) Update(book models.Book, err error) (models.Book, error) {
	return r.Repo.Book.Update(book)
}

func (r *Book) Delete(s string, err error) (string, error) {
	return r.Repo.Book.Delete(s)
}

func NewBook(repo *storage.Storage) *Book {
	return &Book{Repo: repo}
}

type IBookService interface {
	List() ([]models.Book, error)
	Retrieve(string, error) (models.Book, error)
	Create(models.Book, error) (models.Book, error)
	Update(models.Book, error) (models.Book, error)
	Delete(string, error) (string, error)
}
