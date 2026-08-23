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
	rep := repositories.TodoRepository{}
	ser := services.NewTodoService(rep, rep, rep, rep, rep)
	han := handlers.NewTodoHandler(ser, ser, ser, ser)

	db.Init()
	mux := router.SetupRoutes(han)

	mux.Handle("/", http.FileServer(http.Dir("static")))

	http.ListenAndServe(":8080", mux)
}
