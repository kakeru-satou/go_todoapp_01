package main

import (
	"net/http"

	"go-learning/todoApp/db"
	"go-learning/todoApp/router"
)

func main() {

	db.Init()
	mux := router.SetupRoutes()

	mux.Handle("/", http.FileServer(http.Dir("static")))

	http.ListenAndServe(":8080", mux)
}
