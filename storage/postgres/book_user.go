package postgres

import (
	"fmt"
	"gorm.io/gorm"
	"hw1/models"
)

type BookUser struct {
	db *gorm.DB
}

func (r *BookUser) List() ([]models.BookUser, error) {
	sql := "SELECT users.id as user_id, users.username, books.name as book_name, bu.returned, bu.taken_date from users " +
		"LEFT JOIN book_user bu on users.id = bu.user_id " +
		"LEFT JOIN books ON books.name = bu.book_name WHERE bu.returned = true ORDER BY books.name"
	var bookUser []models.BookUser
	tx := r.db.Raw(sql).Scan(&bookUser)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return bookUser, nil
}

func (r *BookUser) Create(bookName string, userID int) (models.BookUser, error) {
	date := "2003.02.03"
	var book models.BookUser
	sql := fmt.Sprintf("INSERT INTO book_user (book_name, user_id, returned, taken_date) VALUES ('%s', %d, true, '%s')", bookName, userID, date)
	tx := r.db.Raw(sql).Scan(&book)
	if tx.Error != nil {
		return models.BookUser{}, tx.Error
	}
	return book, nil
}

func (r *BookUser) ListUserBookCount(date string) ([]models.UserBookCount, error) {
	sql := fmt.Sprintf(
		"SELECT users.username, COUNT(bu.book_name) from users LEFT OUTER JOIN book_user bu on users.id = bu.user_id WHERE bu.taken_date > '%s' GROUP BY users.username;",
		date,
	)
	var userBook []models.UserBookCount
	tx := r.db.Raw(sql).Scan(&userBook)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return userBook, nil
}

func NewBookUser(db *gorm.DB) *BookUser {
	return &BookUser{db: db}
}
