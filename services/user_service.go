package services

import (
	"go-learning/todoApp/models"
	"go-learning/todoApp/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userCreator repositories.UserCreator
}

func NewUserService(userCreator repositories.UserCreator) *UserService {
	return &UserService{
		userCreator: userCreator,
	}
}

func (s *UserService) Signup(name, email, password string) (models.User, error) {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPass),
	}

	return s.userCreator.Create(user)
}
