package handlers

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"strings"

	"go-learning/todoApp/models"
)

type Lister interface {
	GetTodos() ([]models.Todo, error)
}
type Creator interface {
	CreateTodo(task string) (models.Todo, error)
}
type Updater interface {
	UpdateTodo(id string, patch models.PatchTodoRequest) (models.Todo, error)
}
type Deleter interface {
	DeleteTodo(id string) error
}

type TodoHandler struct {
	lister  Lister
	creator Creator
	updater Updater
	deleter Deleter
}

func NewTodoHandler(lister Lister, creator Creator, updater Updater, deleter Deleter) *TodoHandler {
	return &TodoHandler{
		lister:  lister,
		creator: creator,
		updater: updater,
		deleter: deleter,
	}
}

func getIDFromPath(path string) string {

	return strings.TrimPrefix(path, "/api/todoList/")
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, nil)
}

func (h *TodoHandler) GetTodosHandler(w http.ResponseWriter, _ *http.Request) {

	todos, err := h.lister.GetTodos()

	if err != nil {
		http.Error(w, "failed to get todos", http.StatusNotFound)
		return
	}

	err = writeJSON(w, http.StatusOK, todos)

	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *TodoHandler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {

	var req models.CreateTodoRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	todo, err := h.creator.CreateTodo(req.Task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = writeJSON(w, http.StatusCreated, todo)

	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

// PatchTodoHandlerに統合したため使用停止。学習履歴として残置。
// func ToggleTodoHandler(w http.ResponseWriter, r *http.Request) {
//
// 	id := getIDFromPath(r.URL.Path)
//
// 	todo, err := services.ToggleTodo(id)
//
// 	if err != nil {
// 		http.Error(w, "todo not found", http.StatusNotFound)
// 		return
// 	}
//
// 	err = writeJSON(w, http.StatusOK, todo)
//
// 	if err != nil {
// 		http.Error(w, "failed to encode response", http.StatusInternalServerError)
// 	}
// }

func (h *TodoHandler) PatchTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := getIDFromPath(r.URL.Path)

	var req models.PatchTodoRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	todo, err := h.updater.UpdateTodo(id, req)

	if err != nil {
		if errors.Is(err, models.ErrTaskRequired) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "todo not found", http.StatusNotFound)
		return
	}

	err = writeJSON(w, http.StatusOK, todo)

	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *TodoHandler) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {

	id := getIDFromPath(r.URL.Path)

	err := h.deleter.DeleteTodo(id)

	if err != nil {
		http.Error(w, "delete failed", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
