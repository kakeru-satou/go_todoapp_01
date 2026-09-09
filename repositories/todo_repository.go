package repositories

import (
	"go-learning/todoApp/db"
	"go-learning/todoApp/models"
)

type AllFinder interface {
	FindAllTodos(int) ([]models.Todo, error)
}
type Creator interface {
	Create(models.Todo) (models.Todo, error)
}
type Finder interface {
	FindByID(string, int) (models.Todo, error)
}
type Saver interface {
	Save(models.Todo) (models.Todo, error)
}
type Deleter interface{ Delete(models.Todo) error }

type TodoRepository struct{}

func (r TodoRepository) FindAllTodos(userID int) ([]models.Todo, error) {

	var todos []models.Todo
	result := db.DB.Where("user_id = ?", userID).Find(&todos)

	return todos, result.Error
}

func (r TodoRepository) Create(todo models.Todo) (models.Todo, error) {

	result := db.DB.Create(&todo)

	return todo, result.Error
}

func (r TodoRepository) FindByID(todoID string, userID int) (models.Todo, error) {

	var todo models.Todo

	result := db.DB.Where("user_id = ?", userID).First(&todo, todoID)

	return todo, result.Error
}

func (r TodoRepository) Save(todo models.Todo) (models.Todo, error) {

	result := db.DB.Save(&todo)

	return todo, result.Error
}

func (r TodoRepository) Delete(todo models.Todo) error {

	result := db.DB.Delete(&todo)

	return result.Error
}
