package router

import (
	"net/http"

	"go-learning/todoApp/handlers"
)

func SetupRoutes() {

	http.HandleFunc("/api/todoList", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			handlers.GetTodos(w, r)

		case http.MethodPost:
			handlers.CreateTodo(w, r)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/todoList/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPut:
			handlers.ToggleTodo(w, r)

		case http.MethodDelete:
			handlers.DeleteTodo(w, r)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
