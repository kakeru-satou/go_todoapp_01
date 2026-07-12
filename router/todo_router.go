package router

import (
	"net/http"

	"go-learning/todoApp/handlers"
)

func SetupRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/todoList", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			handlers.GetTodosHandler(w, r)

		case http.MethodPost:
			handlers.CreateTodoHandler(w, r)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/todoList/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPut:
			handlers.ToggleTodoHandler(w, r)

		case http.MethodDelete:
			handlers.DeleteTodoHandler(w, r)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
