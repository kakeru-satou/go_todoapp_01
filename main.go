package main

import (
	"net/http"

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

	db.Init()
	todoMux := router.SetupRoutes(todoHan)
	userMux := router.SetupUserRoutes(userHan)

	// Chapter3でミドルウェアの動作確認用に追加したルート。
	// appMuxはtodoMuxを/api/todoList(/)経由でしかマウントしないため、現在は外部から到達不能。
	// 学習履歴として残置。
	// middleHan := middleware.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	claims := r.Context().Value(middleware.UserContextKey).(auth.AccessTokenClaims)
	//
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	json.NewEncoder(w).Encode(claims)
	// }))
	// todoMux.Handle("/api/test/protected", middleHan)

	appMux := http.NewServeMux()

	protectedTodo := middleware.Middleware(todoMux)
	appMux.Handle("/api/todoList", protectedTodo)
	appMux.Handle("/api/todoList/", protectedTodo)

	appMux.Handle("/api/auth/", userMux)
	appMux.Handle("/", http.FileServer(http.Dir("static")))

	http.ListenAndServe(":8080", appMux)
}
