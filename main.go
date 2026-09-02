package main

import (
	"encoding/json"
	"net/http"

	"go-learning/todoApp/auth"
	"go-learning/todoApp/db"
	"go-learning/todoApp/handlers"
	"go-learning/todoApp/middleware"
	"go-learning/todoApp/repositories"
	"go-learning/todoApp/router"
	"go-learning/todoApp/services"
)

func main() {
	todoRep := repositories.TodoRepository{}
	todoSer := services.NewTodoService(todoRep, todoRep, todoRep, todoRep, todoRep)
	todoHan := handlers.NewTodoHandler(todoSer, todoSer, todoSer, todoSer)

	userRep := repositories.UserRepository{}
	userSer := services.NewUserService(userRep, userRep)
	userHan := handlers.NewUserHandler(userSer, userSer)

	middleHan := middleware.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value(middleware.UserContextKey).(auth.AccessTokenClaims)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(claims)
	}))

	db.Init()
	todoMux := router.SetupRoutes(todoHan)
	userMux := router.SetupUserRoutes(userHan)

	todoMux.Handle("/api/auth/", userMux)
	todoMux.Handle("/api/test/protected", middleHan)
	todoMux.Handle("/", http.FileServer(http.Dir("static")))

	http.ListenAndServe(":8080", todoMux)
}
