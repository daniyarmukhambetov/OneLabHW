package in_memory

import (
	"hw1/internal/models"
)

var data []models.User

func init() {
	data = make([]models.User, 0)
}
