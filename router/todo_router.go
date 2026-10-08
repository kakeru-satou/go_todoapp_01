package router

import (
	"go-learning/todoApp/response"
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
			response.WriteErrMessage(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	})

	mux.HandleFunc("/api/todoList/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPatch:
			h.PatchTodoHandler(w, r)

		case http.MethodDelete:
			h.DeleteTodoHandler(w, r)

		default:
			response.WriteErrMessage(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	})

	return mux
}
