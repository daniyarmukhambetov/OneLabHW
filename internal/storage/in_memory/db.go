package in_memory

import (
	"hw1/internal/models"
)

var data []models.UserModel

func init() {
	data = make([]models.UserModel, 0)
}
