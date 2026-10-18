package repositories

import (
	"errors"
	"go-learning/todoApp/db"
	"go-learning/todoApp/models"

	"gorm.io/gorm"
)

type UserCreator interface {
	Create(models.User) (models.User, error)
}

type UserFinder interface {
	GetByEmail(email string) (models.User, error)
	GetByID(id int) (models.User, error)
}

type UserRepository struct{}

func (r UserRepository) Create(user models.User) (models.User, error) {
	result := db.DB.Create(&user)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return user, models.ErrEmailDuplicate
	}

	return user, result.Error
}

func (r UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User

	result := db.DB.First(&user, "email = ?", email)

	return user, result.Error
}

func (r UserRepository) GetByID(id int) (models.User, error) {
	var user models.User

	result := db.DB.First(&user, id)

	return user, result.Error
}
