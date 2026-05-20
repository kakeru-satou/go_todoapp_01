package services

import (
	"go-learning/todoApp/db"
	"go-learning/todoApp/models"
	"go-learning/todoApp/repositories"
)

func GetTodos() ([]models.Todo, error) {

	var todos []models.Todo

	return repositories.FindAllTodos(todos)
}

func CreateTodo(task string) (models.Todo, error) {

	todo := models.Todo{
		Task:   task,
		IsDone: false,
	}

	return repositories.Create(todo)
}

func ToggleTodo(id string) (models.Todo, error) {

	var todo models.Todo

	todo, err := repositories.FindByID(todo, id)

	if err != nil {
		return todo, err
	}

	todo.IsDone = !todo.IsDone

	result := db.DB.Save(&todo)

	return todo, result.Error
}

func DeleteTodo(id string) error {

	var todo models.Todo

	todo, _ = repositories.FindByID(todo, id)

	return repositories.Delete(todo)
}
