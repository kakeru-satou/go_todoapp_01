package services

import (
	"errors"
	"go-learning/todoApp/models"
	"strconv"
)

type FakeRepository struct {
	todos  map[int]models.Todo
	nextID int
}

func NewFakeRepository() *FakeRepository {
	return &FakeRepository{
		todos:  make(map[int]models.Todo),
		nextID: 1,
	}
}

func (f *FakeRepository) FindAllTodos() ([]models.Todo, error) {
	var todoList []models.Todo

	for _, todo := range f.todos {
		todoList = append(todoList, todo)
	}

	return todoList, nil
}

func (f *FakeRepository) Create(todo models.Todo) (models.Todo, error) {
	todo.ID = f.nextID

	f.todos[f.nextID] = todo

	f.nextID = f.nextID + 1

	return todo, nil
}

func (f *FakeRepository) FindByID(findID string) (models.Todo, error) {
	id, err := strconv.Atoi(findID)

	if err != nil {
		return models.Todo{}, errors.New("invalid ID")
	}

	findTodo, ok := f.todos[id]

	if !ok {
		return findTodo, errors.New("not found")
	}

	return findTodo, err
}

func (f *FakeRepository) Save(todo models.Todo) (models.Todo, error) {
	f.todos[todo.ID] = todo

	return todo, nil
}

func (f *FakeRepository) Delete(todo models.Todo) error {
	delete(f.todos, todo.ID)

	return nil
}
