package services

import (
	"go-learning/todoApp/auth"
	"go-learning/todoApp/models"
	"go-learning/todoApp/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userCreator repositories.UserCreator
	userFinder  repositories.UserFinder
}

func NewUserService(userCreator repositories.UserCreator, userFinder repositories.UserFinder) *UserService {
	return &UserService{
		userCreator: userCreator,
		userFinder:  userFinder,
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

func (s *UserService) Signin(email, password string) (string, error) {

	user, err := s.userFinder.GetByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	return auth.GenerateAccessToken(user.ID, user.Name, user.Email)
}
