package postgres

import (
	"fmt"
	"gorm.io/gorm"
	"hw1/internal/models"
)

type BookUser struct {
	db *gorm.DB
}

func (r *BookUser) List() ([]models.BookUser, error) {
	sql := "SELECT users.id as user_id, users.username, books.name as book_name, bu.returned, bu.taken_date from users " +
		"LEFT JOIN book_user bu on users.id = bu.user_id " +
		"LEFT JOIN books ON books.name = bu.book_name WHERE bu.returned = true ORDER BY users.id"
	var bookUser []models.BookUser
	tx := r.db.Raw(sql).Scan(&bookUser)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return bookUser, nil
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
