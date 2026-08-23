package router

import (
	"encoding/json"
	"go-learning/todoApp/db"
	"go-learning/todoApp/handlers"
	"go-learning/todoApp/models"
	"go-learning/todoApp/repositories"
	"go-learning/todoApp/services"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func Setup() (*http.ServeMux, *services.TodoService) {
	db.Init()

	rep := repositories.TodoRepository{}
	ser := services.NewTodoService(rep, rep, rep, rep, rep)
	han := handlers.NewTodoHandler(ser, ser, ser, ser)

	return SetupRoutes(han), ser
}

func SendRequest(mux *http.ServeMux, method string, path string, body ...string) *httptest.ResponseRecorder {

	var reader io.Reader

	if len(body) > 0 {
		reader = strings.NewReader(body[0])
	}

	req := httptest.NewRequest(
		method,
		path,
		reader,
	)

	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	return w
}

func TestRouterGet_Success(t *testing.T) {

	mux, ser := Setup()

	_, err := ser.CreateTodo("test")

	if err != nil {
		t.Fatalf("エラー: %v", err)
	}

	w := SendRequest(mux, http.MethodGet, "/api/todoList")

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
		t.Errorf(
			"想定タスク: %s, 取得タスク: %s",
			"test",
			response[0].Task,
		)
	}
}

func TestRouterPost_Success(t *testing.T) {

	mux, _ := Setup()

	w := SendRequest(mux, http.MethodPost, "/api/todoList", `{"task":"test"}`)

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

func TestRouterPatch_IsDoneSuccess(t *testing.T) {

	mux, ser := Setup()

	todo, err := ser.CreateTodo("test")

	if err != nil {
		t.Fatalf("エラー: %v", err)
	}

	id := strconv.Itoa(todo.ID)

	w := SendRequest(mux, http.MethodPatch, "/api/todoList/"+id, `{"isDone":true}`)

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

func TestRouterPatch_TaskSuccess(t *testing.T) {
	mux, ser := Setup()

	todo, err := ser.CreateTodo("test")

	if err != nil {
		t.Fatalf("エラー: %v", err)
	}

	id := strconv.Itoa(todo.ID)

	w := SendRequest(mux, http.MethodPatch, "/api/todoList/"+id, `{"task":"Test"}`)

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

	if response.Task != "Test" {
		t.Errorf(
			"想定タスク: %s, 取得タスク: %s",
			"Test",
			response.Task,
		)
	}

	if response.IsDone != false {
		t.Errorf(
			"想定状態: %t, 取得状態: %t",
			false,
			response.IsDone,
		)
	}
}

func TestRouterDelete_Success(t *testing.T) {

	mux, ser := Setup()

	repo := repositories.TodoRepository{}

	todo, err := ser.CreateTodo("test")

	if err != nil {
		t.Fatalf("エラー: %v", err)
	}

	id := strconv.Itoa(todo.ID)

	w := SendRequest(mux, http.MethodDelete, "/api/todoList/"+id)

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

	_, err = repo.FindByID(id)

	if err == nil {
		t.Errorf("タスクの削除に失敗")
	}
}

func TestRouterMethodNotAllowed_Success(t *testing.T) {
	mux, _ := Setup()

	w := SendRequest(mux, http.MethodPatch, "/api/todoList")

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
	mux, _ := Setup()

	w := SendRequest(mux, http.MethodGet, "/api/unknown")

	if w.Code != http.StatusNotFound {
		t.Errorf(
			"想定ステータス: %d , 取得ステータス: %d",
			http.StatusNotFound,
			w.Code,
		)
	}
}

func TestRouterPost_EmptyTask(t *testing.T) {
	mux, _ := Setup()

	w := SendRequest(mux, http.MethodPost, "/api/todoList", `{"task":""}`)

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

func TestRouterPatch_NotFound(t *testing.T) {
	mux, _ := Setup()

	w := SendRequest(mux, http.MethodPatch, "/api/todoList/99999", `{"task":"Test"}`)

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

func TestRouterPatch_EmptyTask(t *testing.T) {
	mux, ser := Setup()

	todo, err := ser.CreateTodo("test")

	if err != nil {
		t.Fatalf("エラー: %v", err)
	}

	id := strconv.Itoa(todo.ID)

	w := SendRequest(mux, http.MethodPatch, "/api/todoList/"+id, `{"task":""}`)

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

func TestRouterDelete_NotFound(t *testing.T) {
	mux, _ := Setup()

	w := SendRequest(mux, http.MethodDelete, "/api/todoList/99999")

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

func TestTableRouter(t *testing.T) {
	tests := []struct {
		name   string
		method string
		path   string
		status int
	}{
		{"RouterGet_NotFound", http.MethodGet, "/api/unknown", http.StatusNotFound},
		{"RouterDelete_NotFound", http.MethodDelete, "/api/todoList/9999", http.StatusNotFound},
		{"RouterMethodNotAllowed", http.MethodPatch, "/api/todoList", http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux, _ := Setup()

			w := SendRequest(mux, tt.method, tt.path)

			if w.Code != tt.status {
				t.Errorf(
					"%s: 想定=%d, 実際=%d",
					tt.name,
					tt.status,
					w.Code,
				)
			}
		})
	}
}
