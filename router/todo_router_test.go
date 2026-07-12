package router

import (
	"bytes"
	"encoding/json"
	"go-learning/todoApp/db"
	"go-learning/todoApp/models"
	"go-learning/todoApp/repositories"
	"go-learning/todoApp/services"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func Setup() *http.ServeMux {
	db.Init()

	return SetupRoutes()
}

func TestRouterGet_Success(t *testing.T) {

	mux := Setup()

	_, err := services.CreateTodo("test")

	if err != nil {
		t.Fatalf("エラー: %v", err)
	}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/todoList",
		nil,
	)

	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf(
			"想定ステータス: %d , 取得ステータス: %d",
			http.StatusOK,
			w.Code,
		)
	}

	var response []models.Todo

	err = json.Unmarshal(w.Body.Bytes(), &response)

	if err != nil {
		t.Fatalf("response body: %v", err)
	}

	if len(response) == 0 {
		t.Fatalf("取得件数: %d", len(response))
	}

	if response[0].Task != "test" {
		t.Fatalf(
			"想定タスク: %s, 取得タスク: %s",
			"test",
			response[0].Task,
		)
	}
}

func TestRouterPost_Success(t *testing.T) {

	mux := Setup()

	jsonStr := `{"task":"test"}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/todoList",
		bytes.NewBuffer([]byte(jsonStr)),
	)

	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf(
			"想定ステータス: %d , 取得ステータス: %d",
			http.StatusCreated,
			w.Code,
		)
	}

	var response models.Todo

	err := json.Unmarshal(w.Body.Bytes(), &response)

	if err != nil {
		t.Fatalf("response body: %v", err)
	}

	if response.Task != "test" {
		t.Errorf(
			"想定タスク: %s, 取得タスク: %s",
			"test",
			response.Task,
		)
	}
}

func TestRouterPut_Success(t *testing.T) {

	mux := Setup()

	todo, err := services.CreateTodo("test")

	if err != nil {
		t.Fatalf("エラー: %v", err)
	}

	id := strconv.Itoa(todo.ID)

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/todoList/"+id,
		nil,
	)

	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf(
			"想定ステータス: %d , 取得ステータス: %d",
			http.StatusOK,
			w.Code,
		)
	}

	var response models.Todo

	err = json.Unmarshal(w.Body.Bytes(), &response)

	if err != nil {
		t.Fatalf("response body: %v", err)
	}

	if response.Task != "test" {
		t.Errorf(
			"想定タスク: %s, 取得タスク: %s",
			"test",
			response.Task,
		)
	}

	if response.IsDone != true {
		t.Errorf(
			"想定状態: %t, 取得状態: %t",
			true,
			response.IsDone,
		)
	}
}

func TestRouterDelete_Success(t *testing.T) {

	mux := Setup()

	todo, err := services.CreateTodo("test")

	if err != nil {
		t.Fatalf("エラー: %v", err)
	}

	id := strconv.Itoa(todo.ID)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/todoList/"+id,
		nil,
	)

	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf(
			"想定ステータス: %d , 取得ステータス: %d",
			http.StatusNoContent,
			w.Code,
		)
	}

	if len(w.Body.Bytes()) != 0 {
		t.Errorf(
			"想定した長さ: %d, 取得した長さ: %d",
			0,
			len(w.Body.Bytes()),
		)
	}

	_, err = repositories.FindByID(id)

	if err == nil {
		t.Errorf("タスクの削除に失敗")
	}
}

func TestRouterMethodNotAllowed_Success(t *testing.T) {
	mux := Setup()

	req := httptest.NewRequest(
		http.MethodPatch,
		"/api/todoList",
		nil,
	)

	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf(
			"想定ステータス: %d , 取得ステータス: %d",
			http.StatusMethodNotAllowed,
			w.Code,
		)
	}

	expected := "method not allowed\n"

	if w.Body.String() != expected {
		t.Errorf(
			"想定ボディ: %s , 取得ボディ: %q",
			expected,
			w.Body.String(),
		)
	}
}

func TestRouterGet_NotFound(t *testing.T) {
	mux := Setup()

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/unknown",
		nil,
	)

	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf(
			"想定ステータス: %d , 取得ステータス: %d",
			http.StatusNotFound,
			w.Code,
		)
	}
}

func TestRouterPost_EmptyTask(t *testing.T) {
	mux := Setup()

	jsonStr := `{"task":""}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/todoList",
		bytes.NewBuffer([]byte(jsonStr)),
	)

	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf(
			"想定ステータス: %d , 取得ステータス: %d",
			http.StatusBadRequest,
			w.Code,
		)
	}

	expected := "task is required\n"

	if w.Body.String() != expected {
		t.Errorf(
			"想定ボディ: %s , 取得ボディ: %q",
			expected,
			w.Body.String(),
		)
	}
}

func TestRouterPut_NotFound(t *testing.T) {
	mux := Setup()

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/todoList/99999",
		nil,
	)

	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf(
			"想定ステータス: %d , 取得ステータス: %d",
			http.StatusNotFound,
			w.Code,
		)
	}

	expected := "todo not found\n"

	if w.Body.String() != expected {
		t.Errorf(
			"想定ボディ: %s , 取得ボディ: %q",
			expected,
			w.Body.String(),
		)
	}
}

func TestRouterDelete_NotFound(t *testing.T) {
	mux := Setup()

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/todoList/99999",
		nil,
	)

	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf(
			"想定ステータス: %d , 取得ステータス: %d",
			http.StatusNotFound,
			w.Code,
		)
	}

	expected := "delete failed\n"

	if w.Body.String() != expected {
		t.Errorf(
			"想定ボディ: %s , 取得ボディ: %q",
			expected,
			w.Body.String(),
		)
	}
}
