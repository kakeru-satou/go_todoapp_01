package repositories

import (
	"go-learning/todoApp/db"
	"go-learning/todoApp/models"
	"strconv"
	"testing"
)

func TestCreate_Success(t *testing.T) {

	db.Init()
	todo := models.Todo{
		Task: "test",
	}

	createdTodo, createdErr := Create(todo)
	if createdErr != nil {
		t.Errorf("エラー: %v", createdErr)
	}

	foundTodo, foundErr := FindByID(strconv.Itoa(createdTodo.ID))
	if foundErr != nil {
		t.Errorf("作成失敗")
	}
	if foundTodo.Task != todo.Task {
		t.Errorf("入力内容: %s, 登録内容: %s", foundTodo.Task, todo.Task)
	}
}

func TestFindByID_UnknownID(t *testing.T) {

	db.Init()
	todo := models.Todo{
		Task: "test",
	}

	_, createdErr := Create(todo)
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}

	_, err := FindByID("hoge")
	if err == nil {
		t.Errorf("存在しないタスクの検索")
	}
}

func TestFindByID_Success(t *testing.T) {

	db.Init()
	todo := models.Todo{
		Task: "test",
	}

	createdTodo, createdErr := Create(todo)
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}

	foundTodo, err := FindByID(strconv.Itoa(createdTodo.ID))
	if err != nil {
		t.Errorf("存在しないタスクの検索")
	}
	if foundTodo.Task != createdTodo.Task {
		t.Errorf("タスクが正しく作成できてない")
	}
}

func TestSave_Success(t *testing.T) {

	db.Init()
	todo := models.Todo{
		Task: "test",
	}

	createdTodo, createdErr := Create(todo)
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}
	if createdTodo.IsDone != false {
		t.Fatalf("初期完了状態が不正")
	}

	createdTodo.IsDone = true

	_, savedErr := Save(createdTodo)
	if savedErr != nil {
		t.Errorf("エラー: %v", savedErr)
	}

	foundTodo, foundErr := FindByID(strconv.Itoa(createdTodo.ID))
	if foundErr != nil {
		t.Fatalf("エラー: %v", foundErr)
	}
	if foundTodo.IsDone != createdTodo.IsDone {
		t.Errorf("完了状態の切り替えに失敗")
	}
}

func TestDelete_Success(t *testing.T) {

	db.Init()
	todo := models.Todo{
		Task: "test",
	}

	createdTodo, createdErr := Create(todo)
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}

	deletedErr := Delete(createdTodo)
	if deletedErr != nil {
		t.Errorf("エラー: %v", deletedErr)
	}

	_, foundErr := FindByID(strconv.Itoa(createdTodo.ID))
	if foundErr == nil {
		t.Errorf("削除に失敗")
	}
}
