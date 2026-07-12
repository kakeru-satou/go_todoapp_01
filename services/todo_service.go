package services

import (
	"errors"
	"go-learning/todoApp/models"
	"go-learning/todoApp/repositories"
)

func GetTodos() ([]models.Todo, error) {

	return repositories.FindAllTodos()
}

func CreateTodo(task string) (models.Todo, error) {

	if task == "" {
		return models.Todo{}, errors.New("task is required")
	}

	todo := models.Todo{
		Task:   task,
		IsDone: false,
	}

	return repositories.Create(todo)
}

func ToggleTodo(id string) (models.Todo, error) {

	var todo models.Todo

	todo, err := repositories.FindByID(id)

	if err != nil {
		return todo, err
	}

	todo.IsDone = !todo.IsDone

	return repositories.Save(todo)
}

func DeleteTodo(id string) error {

	var todo models.Todo

	todo, err := repositories.FindByID(id)

	if err != nil {
		return err
	}

	return repositories.Delete(todo)
}
