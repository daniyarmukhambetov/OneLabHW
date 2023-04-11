package postgres

import (
	"fmt"
	"gorm.io/gorm"
	"hw1/internal/models"
)

type User struct {
	db *gorm.DB
}

func (r *User) List() ([]models.User, error) {
	var users []models.User
	tx := r.db.Raw("SELECT * FROM users").Scan(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (r *User) Retrieve(i string) (models.User, error) {
	sql := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", i)
	var user models.User
	tx := r.db.Raw(sql).Scan(&user)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return user, nil
}

func (r *User) Create(in models.UserModelIn) (models.User, error) {
	sql := fmt.Sprintf(
		"INSERT INTO users (username, email, name, last_name, password) VALUES ('%s', '%s', '%s', '%s', '%s')",
		in.Username,
		in.Email,
		in.Name,
		in.LastName,
		in.Password,
	)
	var user models.User
	var id int
	fmt.Println(id)
	tx := r.db.Raw(sql).Scan(&id)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return user, nil
}

func (r *User) Update(i string, in models.UserUpdate) (models.User, error) {
	sql := fmt.Sprintf("UPDATE users SET name = '%s', last_name = '%s', password = '%s' WHERE username = '%s'", in.Name, in.LastName, in.Password, i)
	var user models.User
	tx := r.db.Raw(sql).Scan(&user)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return user, nil
}

func (r *User) Delete(i string) (string, error) {
	sql := fmt.Sprintf("DELETE FROM users WHERE username = %s", i)
	tx := r.db.Raw(sql).Scan(&i)
	if tx.Error != nil {
		return "", tx.Error
	}
	return i, nil
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}
