package services

import (
	"errors"
	"go-learning/todoApp/models"
	"strconv"
)

type FakeTodoRepository struct {
	todos  map[int]models.Todo
	nextID int
	userID int
}

func NewFakeTodoRepository() *FakeTodoRepository {
	return &FakeTodoRepository{
		todos:  make(map[int]models.Todo),
		nextID: 1,
		userID: 0,
	}
}

func (f *FakeTodoRepository) FindAllTodos(userID int) ([]models.Todo, error) {
	var todoList []models.Todo

	for _, todo := range f.todos {
		if userID == todo.UserID {
			todoList = append(todoList, todo)
		}
	}

	return todoList, nil
}

func (f *FakeTodoRepository) Create(todo models.Todo) (models.Todo, error) {
	todo.TodoID = f.nextID

	f.todos[f.nextID] = todo

	f.nextID = f.nextID + 1

	return todo, nil
}

func (f *FakeTodoRepository) FindByID(findID string, userID int) (models.Todo, error) {
	findTodoID, err := strconv.Atoi(findID)

	if err != nil {
		return models.Todo{}, errors.New("invalid ID")
	}

	findTodo, ok := f.todos[findTodoID]

	if !ok || findTodo.UserID != userID {
		return findTodo, errors.New("not found")
	}

	return findTodo, err
}

func (f *FakeTodoRepository) Save(todo models.Todo) (models.Todo, error) {
	f.todos[todo.TodoID] = todo

	return todo, nil
}

func (f *FakeTodoRepository) Delete(todo models.Todo) error {
	delete(f.todos, todo.TodoID)

	return nil
}
