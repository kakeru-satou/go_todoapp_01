package handlers

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strings"

	"go-learning/todoApp/auth"
	"go-learning/todoApp/middleware"
	"go-learning/todoApp/models"
	"go-learning/todoApp/response"
)

type Lister interface {
	GetTodos(userID int) ([]models.Todo, error)
}
type Creator interface {
	CreateTodo(task string, userID int) (models.Todo, error)
}
type Updater interface {
	UpdateTodo(todoID string, userID int, patch models.PatchTodoRequest) (models.Todo, error)
}
type Deleter interface {
	DeleteTodo(todoID string, userID int) error
}

type TodoHandler struct {
	lister  Lister
	creator Creator
	updater Updater
	deleter Deleter
}

func NewTodoHandler(lister Lister, creator Creator, updater Updater, deleter Deleter) *TodoHandler {
	return &TodoHandler{
		lister:  lister,
		creator: creator,
		updater: updater,
		deleter: deleter,
	}
}

func getIDFromPath(path string) string {
	return strings.TrimPrefix(path, "/api/todoList/")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, nil)
}

func (h *TodoHandler) GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserContextKey).(auth.AccessTokenClaims)
	todos, err := h.lister.GetTodos(claims.UserID)

	if err != nil {
		response.WriteErrMessage(w, http.StatusNotFound, "failed to get todos")
		return
	}

	err = response.WriteJSON(w, http.StatusOK, todos)

	if err != nil {
		response.WriteErrMessage(w, http.StatusInternalServerError, "failed to encode response")
	}
}

func (h *TodoHandler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTodoRequest
	claims := r.Context().Value(middleware.UserContextKey).(auth.AccessTokenClaims)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteErrMessage(w, http.StatusBadRequest, "invalid request")
		return
	}

	todo, err := h.creator.CreateTodo(req.Task, claims.UserID)

	if err != nil {
		if errors.Is(err, models.ErrTaskRequired) {
			response.WriteErrMessage(w, http.StatusBadRequest, "タスクが空です")
			return
		}
		log.Println(err)
		response.WriteErrMessage(w, http.StatusInternalServerError, "タスクの作成に失敗しました")
		return
	}

	err = response.WriteJSON(w, http.StatusCreated, todo)

	if err != nil {
		response.WriteErrMessage(w, http.StatusInternalServerError, "failed to encode response")
	}
}

// PatchTodoHandlerに統合したため使用停止。学習履歴として残置。
// func ToggleTodoHandler(w http.ResponseWriter, r *http.Request) {
//
// 	id := getIDFromPath(r.URL.Path)
//
// 	todo, err := services.ToggleTodo(id)
//
// 	if err != nil {
// 		http.Error(w, "todo not found", http.StatusNotFound)
// 		return
// 	}
//
// 	err = writeJSON(w, http.StatusOK, todo)
//
// 	if err != nil {
// 		http.Error(w, "failed to encode response", http.StatusInternalServerError)
// 	}
// }

func (h *TodoHandler) PatchTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := getIDFromPath(r.URL.Path)

	var req models.PatchTodoRequest
	claims := r.Context().Value(middleware.UserContextKey).(auth.AccessTokenClaims)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteErrMessage(w, http.StatusBadRequest, "invalid request")
		return
	}

	todo, err := h.updater.UpdateTodo(todoID, claims.UserID, req)

	if err != nil {
		if errors.Is(err, models.ErrTaskRequired) {
			response.WriteErrMessage(w, http.StatusBadRequest, "タスクが空です")
			return
		}
		response.WriteErrMessage(w, http.StatusNotFound, "todo not found")
		return
	}

	err = response.WriteJSON(w, http.StatusOK, todo)

	if err != nil {
		response.WriteErrMessage(w, http.StatusInternalServerError, "failed to encode response")
	}
}

func (h *TodoHandler) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := getIDFromPath(r.URL.Path)
	claims := r.Context().Value(middleware.UserContextKey).(auth.AccessTokenClaims)

	err := h.deleter.DeleteTodo(todoID, claims.UserID)

	if err != nil {
		response.WriteErrMessage(w, http.StatusNotFound, "削除に失敗しました")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
