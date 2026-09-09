package repositories

import (
	"go-learning/todoApp/db"
	"go-learning/todoApp/models"
)

type UserCreator interface {
	Create(models.User) (models.User, error)
}

type UserFinder interface {
	GetByEmail(email string) (models.User, error)
}

type UserRepository struct{}

func (r UserRepository) Create(user models.User) (models.User, error) {

	result := db.DB.Create(&user)

	return user, result.Error
}

func (r UserRepository) GetByEmail(email string) (models.User, error) {

	var user models.User

	result := db.DB.First(&user, "email = ?", email)

	return user, result.Error
}
