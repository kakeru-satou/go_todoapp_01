package services

import (
	"go-learning/todoApp/models"
	"strconv"
	"testing"
)

func Setup() *TodoService {
	rep := NewFakeRepository()
	ser := NewTodoService(rep, rep, rep, rep, rep)

	return ser
}

func SetupMock() (*TodoService, *MockDeleter) {
	rep := NewFakeRepository()
	deleter := MockDeleter{}
	ser := NewTodoService(rep, rep, rep, rep, &deleter)

	return ser, &deleter
}

func TestCreateTodo_EmptyTask(t *testing.T) {

	ser := Setup()
	todo, err := ser.CreateTodo("")

	if err == nil {
		t.Errorf("エラーが返ってない")
	}

	if todo.Task != "" {
		t.Errorf("作成したタスクがnilでない")
	}
}

func TestCreateTodo_Success(t *testing.T) {

	ser := Setup()
	task := "test"

	todo, err := ser.CreateTodo(task)

	if err != nil {
		t.Errorf("エラー: %v", err)
	}

	if todo.Task != task {
		t.Errorf("入力内容: %s, 登録内容: %s", task, todo.Task)
	}
}

// ToggleTodoはUpdateTodo(PATCH化)に統合したため使用停止。学習履歴として残置。
// func TestToggleTodo_UnknownID(t *testing.T) {
//
// 	db.Init()
//
// 	todo, createdErr := CreateTodo("test")
// 	if createdErr != nil {
// 		t.Fatalf("エラー: %v", createdErr)
// 	}
//
// 	if todo.IsDone != false {
// 		t.Fatalf("初期完了状態が不正")
// 	}
//
// 	_, toggleErr := ToggleTodo("hoge")
// 	if toggleErr == nil {
// 		t.Errorf("存在しないタスクを切り替え")
// 	}
// }
//
// func TestToggleTodo_Success(t *testing.T) {
//
// 	db.Init()
//
// 	todo, createdErr := CreateTodo("test")
// 	if createdErr != nil {
// 		t.Fatalf("エラー: %v", createdErr)
// 	}
//
// 	if todo.IsDone != false {
// 		t.Errorf("初期完了状態が不正")
// 	}
//
// 	updatedTodo, toggleErr := ToggleTodo(strconv.Itoa(todo.ID))
// 	if toggleErr != nil {
// 		t.Fatalf("エラー: %v", toggleErr)
// 	}
//
// 	if updatedTodo.IsDone != true {
// 		t.Errorf("完了状態切り替えに失敗")
// 	}
// }

func TestUpdateTodoIsDone_Success(t *testing.T) {

	ser := Setup()

	todo, createdErr := ser.CreateTodo("test")
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}

	isDone := !todo.IsDone
	patchReq := models.PatchTodoRequest{
		IsDone: &isDone,
	}

	updatedTodo, err := ser.UpdateTodo(strconv.Itoa(todo.ID), patchReq)
	if err != nil {
		t.Errorf("エラー: %v", err)
	}
	if updatedTodo.IsDone != isDone {
		t.Errorf("想定完了状態: %v, 取得完了状態: %v", isDone, updatedTodo.IsDone)
	}
	if updatedTodo.Task != todo.Task {
		t.Errorf("想定タスク: %v, 取得タスク: %v", todo.Task, updatedTodo.Task)
	}
}

func TestUpdateTodoTask_Success(t *testing.T) {

	ser := Setup()

	todo, createdErr := ser.CreateTodo("test")
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}

	task := "update"
	patchReq := models.PatchTodoRequest{
		Task: &task,
	}

	updatedTodo, err := ser.UpdateTodo(strconv.Itoa(todo.ID), patchReq)
	if err != nil {
		t.Errorf("エラー: %v", err)
	}
	if updatedTodo.IsDone != todo.IsDone {
		t.Errorf("想定完了状態: %v, 取得完了状態: %v", todo.IsDone, updatedTodo.IsDone)
	}
	if updatedTodo.Task != task {
		t.Errorf("想定タスク: %v, 取得タスク: %v", task, updatedTodo.Task)
	}
}

func TestUpdateTodoTask_UnknownID(t *testing.T) {

	ser := Setup()

	task := "test"
	patchReq := models.PatchTodoRequest{
		Task: &task,
	}

	_, err := ser.UpdateTodo("99999999", patchReq)
	if err == nil {
		t.Errorf("エラーが返っていない")
	}
}

func TestUpdateTodoTask_EmptyTask(t *testing.T) {

	ser := Setup()

	todo, createdErr := ser.CreateTodo("test")
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}

	task := ""
	patchReq := models.PatchTodoRequest{
		Task: &task,
	}

	_, err := ser.UpdateTodo(strconv.Itoa(todo.ID), patchReq)
	if err == nil {
		t.Errorf("エラーが返っていない")
	}
}

func TestDeleteTodo_UnknownID(t *testing.T) {

	ser := Setup()

	_, createdErr := ser.CreateTodo("test")
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}

	deleteErr := ser.DeleteTodo("hoge")
	if deleteErr == nil {
		t.Errorf("存在しないタスクを削除")
	}
}

func TestDeleteTodo_Success(t *testing.T) {

	ser := Setup()

	todo, createdErr := ser.CreateTodo("test")
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}

	deleteErr := ser.DeleteTodo(strconv.Itoa(todo.ID))
	if deleteErr != nil {
		t.Fatalf("エラー: %v", deleteErr)
	}

	allTodos, getError := ser.GetTodos()
	if getError != nil {
		t.Fatalf("エラー: %v", getError)
	}

	for _, getTodo := range allTodos {
		if getTodo.ID == todo.ID {
			t.Errorf("削除に失敗")
		}
	}
}

func TestDeleteTodo_Mock(t *testing.T) {
	ser, m := SetupMock()

	todo, createdErr := ser.CreateTodo("test")
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}

	err := ser.DeleteTodo(strconv.Itoa(todo.ID))
	if err != nil {
		t.Errorf("エラー: %v", err)
	}
	if m.CallCount != 1 {
		t.Errorf("mockが呼ばれていない")
	}
	if m.Todo != todo {
		t.Errorf("mock内のtodoが不一致")
	}
}
