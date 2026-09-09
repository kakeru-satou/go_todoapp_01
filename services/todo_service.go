package services

import (
	"go-learning/todoApp/models"
	"go-learning/todoApp/repositories"
)

type TodoService struct {
	allFinder repositories.AllFinder
	creator   repositories.Creator
	finder    repositories.Finder
	saver     repositories.Saver
	deleter   repositories.Deleter
}

func NewTodoService(allFinder repositories.AllFinder, creator repositories.Creator, finder repositories.Finder, saver repositories.Saver, deleter repositories.Deleter) *TodoService {
	return &TodoService{
		allFinder: allFinder,
		creator:   creator,
		finder:    finder,
		saver:     saver,
		deleter:   deleter,
	}
}

func (s *TodoService) GetTodos(userID int) ([]models.Todo, error) {

	return s.allFinder.FindAllTodos(userID)
}

func (s *TodoService) CreateTodo(task string, userID int) (models.Todo, error) {

	if task == "" {
		return models.Todo{}, models.ErrTaskRequired
	}

	todo := models.Todo{
		Task:   task,
		IsDone: false,
		UserID: userID,
	}

	return s.creator.Create(todo)
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

func (s *TodoService) UpdateTodo(todoID string, userID int, patch models.PatchTodoRequest) (models.Todo, error) {
	todo, err := s.finder.FindByID(todoID, userID)

	if err != nil {
		return todo, err
	}

	if patch.IsDone != nil {
		todo.IsDone = *patch.IsDone
	}
	if patch.Task != nil {
		if *patch.Task == "" {
			return models.Todo{}, models.ErrTaskRequired
		}
		todo.Task = *patch.Task
	}

	return s.saver.Save(todo)
}

func (s *TodoService) DeleteTodo(todoID string, userID int) error {

	var todo models.Todo

	todo, err := s.finder.FindByID(todoID, userID)

	if err != nil {
		return err
	}

	return s.deleter.Delete(todo)
}
