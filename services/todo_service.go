package services

import (
	"errors"
	"go-learning/todoApp/models"
	"go-learning/todoApp/repositories"
)

var ErrTaskRequired = errors.New("task is required")

func GetTodos() ([]models.Todo, error) {

	return repositories.FindAllTodos()
}

func CreateTodo(task string) (models.Todo, error) {

	if task == "" {
		return models.Todo{}, ErrTaskRequired
	}

	todo := models.Todo{
		Task:   task,
		IsDone: false,
	}

	return repositories.Create(todo)
}

// UpdateTodo(PATCH化)に統合したため使用停止。学習履歴として残置。
// func ToggleTodo(id string) (models.Todo, error) {
//
// 	var todo models.Todo
//
// 	todo, err := repositories.FindByID(id)
//
// 	if err != nil {
// 		return todo, err
// 	}
//
// 	todo.IsDone = !todo.IsDone
//
// 	return repositories.Save(todo)
// }

func UpdateTodo(id string, patch models.PatchTodoRequest) (models.Todo, error) {
	todo, err := repositories.FindByID(id)

	if err != nil {
		return todo, err
	}

	if patch.IsDone != nil {
		todo.IsDone = *patch.IsDone
	}
	if patch.Task != nil {
		if *patch.Task == "" {
			return models.Todo{}, ErrTaskRequired
		}
		todo.Task = *patch.Task
	}

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
