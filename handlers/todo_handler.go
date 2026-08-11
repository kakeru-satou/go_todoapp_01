package handlers

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"strings"

	"go-learning/todoApp/models"
	"go-learning/todoApp/services"
)

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

func GetTodosHandler(w http.ResponseWriter, _ *http.Request) {

	todos, err := services.GetTodos()

	if err != nil {
		http.Error(w, "failed to get todos", http.StatusNotFound)
		return
	}

	err = writeJSON(w, http.StatusOK, todos)

	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {

	var req models.CreateTodoRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	todo, err := services.CreateTodo(req.Task)

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

func PatchTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := getIDFromPath(r.URL.Path)

	var req models.PatchTodoRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	todo, err := services.UpdateTodo(id, req)

	if err != nil {
		if errors.Is(err, services.ErrTaskRequired) {
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

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {

	id := getIDFromPath(r.URL.Path)

	err := services.DeleteTodo(id)

	if err != nil {
		http.Error(w, "delete failed", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
