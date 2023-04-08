package postgres

import (
	"fmt"
	"gorm.io/gorm"
	"hw1/internal/models"
)

type Book struct {
	db *gorm.DB
}

func (r *Book) List() ([]models.Book, error) {
	var books []models.Book
	res := r.db.Raw("select * from books;").Scan(&books)
	if res.Error != nil {
		return nil, res.Error
	}
	return books, nil
}

func (r *Book) Retrieve(name string) (models.Book, error) {
	sql := fmt.Sprintf("SELECT * FROM books WHERE name = '%s'", name)
	tx := r.db.Raw(sql)
	if tx.Error != nil {
		return models.Book{}, tx.Error
	}
	var book models.Book
	tx = tx.Scan(&book)
	if tx.Error != nil {
		return models.Book{}, tx.Error
	}
	return book, nil
}

func (r *Book) Create(book models.Book) (models.Book, error) {
	sql := fmt.Sprintf("INSERT INTO books (name, author) VALUES ('%s', '%s')", book.Name, book.Author)
	tx := r.db.Raw(sql)
	if tx.Error != nil {
		return models.Book{}, tx.Error
	}
	tx.Scan(&book)
	if tx.Error != nil {
		return models.Book{}, tx.Error
	}
	return book, nil
}

func (r *Book) Update(book models.Book) (models.Book, error) {
	sql := fmt.Sprintf("UPDATE books SET name = '%s', author = '%s'", book.Name, book.Author)
	tx := r.db.Raw(sql).Scan(&book)
	if tx.Error != nil {
		return models.Book{}, tx.Error
	}
	return book, nil
}

func (r *Book) Delete(name string) (string, error) {
	sql := fmt.Sprintf("DELETE FROM books WHERE name = '%s'", name)
	tx := r.db.Raw(sql).Scan(&name)
	if tx.Error != nil {
		return "", tx.Error
	}
	return name, nil
}

func NewBook(db *gorm.DB) *Book {
	return &Book{db: db}
}
