package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"go-learning/todoApp/db"
	"go-learning/todoApp/models"
	"go-learning/todoApp/services"
)

func TestCreateTodoHandler_Success(t *testing.T) {

	db.Init()

	jsonStr := `{"task":"test"}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/todoList",
		bytes.NewBuffer([]byte(jsonStr)),
	)
	w := httptest.NewRecorder()

	CreateTodoHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf(
			"想定ステータス: %d , 取得したステータス: %d",
			http.StatusCreated,
			w.Code,
		)
	}

	var responseTodo models.Todo

	err := json.Unmarshal(w.Body.Bytes(), &responseTodo)
	if err != nil {
		t.Fatalf("response body: %v", err)
	}

	if responseTodo.Task != "test" {
		t.Errorf(
			"想定タスク: %s, 取得タスク: %s",
			"test",
			responseTodo.Task,
		)
	}
}

func TestCreateTodoHandler_EmptyTask(t *testing.T) {

	db.Init()

	jsonStr := `{"task":""}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/todoList",
		bytes.NewBuffer([]byte(jsonStr)),
	)
	w := httptest.NewRecorder()

	CreateTodoHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf(
			"想定ステータス: %d , 取得したステータス %d",
			http.StatusBadRequest,
			w.Code,
		)
	}
}

func TestGetTodosHandler_Success(t *testing.T) {

	db.Init()

	_, err1 := services.CreateTodo("test1")
	_, err2 := services.CreateTodo("test2")
	if err1 != nil || err2 != nil {
		t.Fatalf("todo作成に失敗")
	}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/todoList",
		nil,
	)

	w := httptest.NewRecorder()

	GetTodosHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf(
			"想定ステータス: %d , 取得したステータス %d",
			http.StatusOK,
			w.Code,
		)
	}
}

// ToggleTodoHandlerはPatchTodoHandlerに統合したため使用停止。学習履歴として残置。
// func TestToggleTodoHandler_Success(t *testing.T) {
//
// 	db.Init()
//
// 	todo, err := services.CreateTodo("test")
// 	if err != nil {
// 		t.Fatalf("todo作成に失敗")
// 	}
//
// 	id := strconv.Itoa(todo.ID)
//
// 	req := httptest.NewRequest(
// 		http.MethodPut,
// 		"/api/todoList/"+id,
// 		nil,
// 	)
//
// 	w := httptest.NewRecorder()
//
// 	ToggleTodoHandler(w, req)
//
// 	if w.Code != http.StatusOK {
// 		t.Errorf(
// 			"想定ステータス: %d , 取得したステータス %d",
// 			http.StatusOK,
// 			w.Code,
// 		)
// 	}
// }
//
// func TestToggleTodoHandler_NotFound(t *testing.T) {
//
// 	db.Init()
//
// 	req := httptest.NewRequest(
// 		http.MethodPut,
// 		"/api/todoList/99999",
// 		nil,
// 	)
//
// 	w := httptest.NewRecorder()
//
// 	ToggleTodoHandler(w, req)
//
// 	if w.Code != http.StatusNotFound {
// 		t.Errorf(
// 			"想定ステータス: %d , 取得したステータス %d",
// 			http.StatusNotFound,
// 			w.Code,
// 		)
// 	}
// }

func TestDeleteTodoHandler_Success(t *testing.T) {

	db.Init()

	todo, err := services.CreateTodo("test")
	if err != nil {
		t.Fatalf("todo作成に失敗")
	}

	id := strconv.Itoa(todo.ID)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/todoList/"+id,
		nil,
	)

	w := httptest.NewRecorder()

	DeleteTodoHandler(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf(
			"想定ステータス: %d , 取得したステータス %d",
			http.StatusNoContent,
			w.Code,
		)
	}
}

func TestDeleteTodoHandler_NotFound(t *testing.T) {

	db.Init()

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/todoList/99999",
		nil,
	)

	w := httptest.NewRecorder()

	DeleteTodoHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf(
			"想定ステータス: %d , 取得したステータス %d",
			http.StatusNotFound,
			w.Code,
		)
	}
}
