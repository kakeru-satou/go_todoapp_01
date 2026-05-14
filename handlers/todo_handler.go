package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"

	"go-learning/todoApp/db"
	"go-learning/todoApp/models"
)

func getIDFromPath(path string) string {
	return strings.TrimPrefix(path, "/api/todoList/")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, nil)
}

func GetTodos(w http.ResponseWriter, _ *http.Request) {
	var todos []models.Todo
	db.DB.Find(&todos)

	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Task string `json:"task"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if req.Task == "" {
		http.Error(w, "task is empty", http.StatusBadRequest)
		return
	}

	todo := models.Todo{
		Task:   req.Task,
		IsDone: false,
	}

	db.DB.Create(&todo)
	json.NewEncoder(w).Encode(todo)
}

func ToggleTodo(w http.ResponseWriter, r *http.Request) {

	id := getIDFromPath(r.URL.Path)

	var todo models.Todo

	if err := db.DB.First(&todo, id).Error; err != nil {
		http.Error(w, "todo not found", http.StatusNotFound)
		return
	}

	todo.IsDone = !todo.IsDone

	db.DB.Save(&todo)

	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

	id := getIDFromPath(r.URL.Path)

	if err := db.DB.Delete(&models.Todo{}, id).Error; err != nil {
		http.Error(w, "delete failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
