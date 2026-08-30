package main

import (
	"net/http"

	"go-learning/todoApp/db"
	"go-learning/todoApp/handlers"
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

	db.Init()
	todoMux := router.SetupRoutes(todoHan)
	userMux := router.SetupUserRoutes(userHan)

	todoMux.Handle("/api/auth/", userMux)
	todoMux.Handle("/", http.FileServer(http.Dir("static")))

	http.ListenAndServe(":8080", todoMux)
}
