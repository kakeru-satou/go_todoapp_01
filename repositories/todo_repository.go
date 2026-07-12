package repositories

import (
	"go-learning/todoApp/db"
	"go-learning/todoApp/models"
)

func FindAllTodos() ([]models.Todo, error) {

	var todos []models.Todo
	result := db.DB.Find(&todos)

	return todos, result.Error
}

func Create(todo models.Todo) (models.Todo, error) {

	result := db.DB.Create(&todo)

	return todo, result.Error
}

func FindByID(id string) (models.Todo, error) {

	var todo models.Todo

	result := db.DB.First(&todo, id)

	return todo, result.Error
}

func Save(todo models.Todo) (models.Todo, error) {

	result := db.DB.Save(&todo)

	return todo, result.Error
}

func Delete(todo models.Todo) error {

	result := db.DB.Delete(&todo)

	return result.Error
}
