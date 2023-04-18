package models

type User struct {
	ID       int `gorm:"primaryKey"`
	Username string
	Email    string
	Name     string
	LastName string
	Password string
}

type UserModelOut struct {
	ID       int
	Username string
	Email    string
	Name     string
	LastName string
}

type UserModelIn struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
}

type UserUpdate struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
}
