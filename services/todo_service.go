package services

import (
	"go-learning/todoApp/db"
	"go-learning/todoApp/models"
)

func GetTodos() ([]models.Todo, error) {

	var todos []models.Todo

	result := db.DB.Find(&todos)

	return todos, result.Error
}

func CreateTodo(task string) (models.Todo, error) {

	todo := models.Todo{
		Task:   task,
		IsDone: false,
	}

	result := db.DB.Create(&todo)

	return todo, result.Error
}

func ToggleTodo(id string) (models.Todo, error) {

	var todo models.Todo

	if err := db.DB.First(&todo, id).Error; err != nil {
		return todo, err
	}

	todo.IsDone = !todo.IsDone

	result := db.DB.Save(&todo)

	return todo, result.Error
}

func DeleteTodo(id string) error {

	return db.DB.Delete(&models.Todo{}, id).Error
}
