package models

type Book struct {
	Name   string `json:"name" gorm:"primaryKey"`
	Author string `json:"author" gorm:"not null"`
}
