package repositories

import (
	"go-learning/todoApp/db"
	"go-learning/todoApp/models"
)

type UserCreator interface {
	Create(models.User) (models.User, error)
}

type UserRepository struct{}

func (r UserRepository) Create(user models.User) (models.User, error) {
	result := db.DB.Create(&user)

	return user, result.Error
}
