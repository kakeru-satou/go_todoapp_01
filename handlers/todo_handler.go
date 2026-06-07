package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"

	"go-learning/todoApp/models"
	"go-learning/todoApp/services"
)

func getIDFromPath(path string) string {

	return strings.TrimPrefix(path, "/api/todoList/")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, nil)
}

func GetTodosHandler(w http.ResponseWriter, _ *http.Request) {

	todos, err := services.GetTodos()

	if err != nil {
		http.Error(w, "failed to get todos", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
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

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(todo)
}

func ToggleTodoHandler(w http.ResponseWriter, r *http.Request) {

	id := getIDFromPath(r.URL.Path)

	todo, err := services.ToggleTodo(id)

	if err != nil {
		http.Error(w, "todo not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {

	id := getIDFromPath(r.URL.Path)

	if err := services.DeleteTodo(id); err != nil {
		http.Error(w, "delete failed", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
