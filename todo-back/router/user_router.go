package router

import (
	"go-learning/todoApp/response"
	"net/http"
)

type UserHandler interface {
	SignupHandler(w http.ResponseWriter, r *http.Request)
	SigninHandler(w http.ResponseWriter, r *http.Request)
	MeHandler(w http.ResponseWriter, r *http.Request)
}

func SetupUserRoutes(h UserHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/auth/signup", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.SignupHandler(w, r)
		default:
			response.WriteErrMessage(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	})

	mux.HandleFunc("/api/auth/signin", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.SigninHandler(w, r)
		default:
			response.WriteErrMessage(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	})

	mux.HandleFunc("/api/users/me", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.MeHandler(w, r)
		default:
			response.WriteErrMessage(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	})

	return mux
}
