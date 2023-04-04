package in_memory

import (
	"hw1/internal/models"
)

var data []models.UserModel

func init() { // использование init bad practise 
	data = make([]models.UserModel, 0) // лучше сделать map потому что так будет гарантия уникальности ID и проще искать по ID 
}
