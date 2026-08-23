package services

import "go-learning/todoApp/models"

type MockDeleter struct {
	CallCount int
	Todo      models.Todo
	Error     error
}

func (m *MockDeleter) Delete(todo models.Todo) error {
	m.CallCount++
	m.Todo = todo
	return m.Error
}
