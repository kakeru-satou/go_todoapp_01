package router

import (
	"net/http"
)

type Handler interface {
	GetTodosHandler(w http.ResponseWriter, r *http.Request)
	CreateTodoHandler(w http.ResponseWriter, r *http.Request)
	PatchTodoHandler(w http.ResponseWriter, r *http.Request)
	DeleteTodoHandler(w http.ResponseWriter, r *http.Request)
}

func SetupRoutes(h Handler) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/todoList", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			h.GetTodosHandler(w, r)

		case http.MethodPost:
			h.CreateTodoHandler(w, r)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/todoList/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPatch:
			h.PatchTodoHandler(w, r)

		case http.MethodDelete:
			h.DeleteTodoHandler(w, r)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
