package services

import (
	"go-learning/todoApp/db"
	"strconv"
	"testing"
)

func TestCreateTodo_EmptyTask(t *testing.T) {

	db.Init()
	todo, err := CreateTodo("")

	if err == nil {
		t.Errorf("エラーが返ってない")
	}

	if todo.Task != "" {
		t.Errorf("作成したタスクがnilでない")
	}
}

func TestCreateTodo_Success(t *testing.T) {

	db.Init()
	task := "test"

	todo, err := CreateTodo(task)

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

func TestDeleteTodo_UnknownID(t *testing.T) {

	db.Init()

	_, createdErr := CreateTodo("test")
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}

	deleteErr := DeleteTodo("hoge")
	if deleteErr == nil {
		t.Errorf("存在しないタスクを削除")
	}
}

func TestDeleteTodo_Success(t *testing.T) {

	db.Init()

	todo, createdErr := CreateTodo("test")
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}

	deleteErr := DeleteTodo(strconv.Itoa(todo.ID))
	if deleteErr != nil {
		t.Fatalf("エラー: %v", deleteErr)
	}

	allTodos, getError := GetTodos()
	if getError != nil {
		t.Fatalf("エラー: %v", getError)
	}

	for _, getTodo := range allTodos {
		if getTodo.ID == todo.ID {
			t.Errorf("削除に失敗")
		}
	}
}
