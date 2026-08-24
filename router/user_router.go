package router

import "net/http"

type UserHandler interface {
	SignupHandler(w http.ResponseWriter, r *http.Request)
}

func SetupUserRoutes(h UserHandler) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/auth/signup", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.SignupHandler(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
