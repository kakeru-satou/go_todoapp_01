package services

import (
	"errors"
	"go-learning/todoApp/models"
)

type FakeUserRepository struct {
	users  map[int]models.User
	nextID int
}

func NewFakeUserRepository() *FakeUserRepository {
	return &FakeUserRepository{
		users:  make(map[int]models.User),
		nextID: 1,
	}
}

func (f *FakeUserRepository) Create(user models.User) (models.User, error) {
	for _, foundUser := range f.users {
		if foundUser.Email == user.Email {
			return user, errors.New("invalid email")
		}
	}

	user.ID = f.nextID

	f.users[f.nextID] = user

	f.nextID++

	return user, nil
}

func (f *FakeUserRepository) GetByEmail(email string) (models.User, error) {
	var foundUser models.User
	isFound := false

	for _, user := range f.users {
		if email == user.Email {
			foundUser = user
			isFound = true
		}
	}
	if !isFound {
		return foundUser, errors.New("not found")
	}

	return foundUser, nil
}
