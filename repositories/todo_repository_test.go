package repositories

import (
	"go-learning/todoApp/db"
	"go-learning/todoApp/models"
	"strconv"
	"testing"
)

func TestCreate_Success(t *testing.T) {

	db.Init()
	repo := TodoRepository{}
	todo := models.Todo{
		Task:   "test",
		UserID: 0,
	}

	createdTodo, createdErr := repo.Create(todo)
	if createdErr != nil {
		t.Errorf("エラー: %v", createdErr)
	}

	foundTodo, foundErr := repo.FindByID(strconv.Itoa(createdTodo.TodoID), createdTodo.UserID)
	if foundErr != nil {
		t.Errorf("作成失敗")
	}
	if foundTodo.Task != todo.Task {
		t.Errorf("入力内容: %s, 登録内容: %s", foundTodo.Task, todo.Task)
	}
}

func TestFindByID_UnknownID(t *testing.T) {

	db.Init()
	repo := TodoRepository{}
	todo := models.Todo{
		Task:   "test",
		UserID: 0,
	}

	createdTodo, createdErr := repo.Create(todo)
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}

	_, err := repo.FindByID("hoge", createdTodo.UserID)
	if err == nil {
		t.Errorf("存在しないタスクの検索")
	}
}

func TestFindByID_Success(t *testing.T) {

	db.Init()
	repo := TodoRepository{}
	todo := models.Todo{
		Task:   "test",
		UserID: 0,
	}

	createdTodo, createdErr := repo.Create(todo)
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}

	foundTodo, err := repo.FindByID(strconv.Itoa(createdTodo.TodoID), createdTodo.UserID)
	if err != nil {
		t.Errorf("存在しないタスクの検索")
	}
	if foundTodo.Task != createdTodo.Task {
		t.Errorf("タスクが正しく作成できてない")
	}
}

func TestSave_Success(t *testing.T) {

	db.Init()
	repo := TodoRepository{}
	todo := models.Todo{
		Task:   "test",
		UserID: 0,
	}

	createdTodo, createdErr := repo.Create(todo)
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}
	if createdTodo.IsDone != false {
		t.Fatalf("初期完了状態が不正")
	}

	createdTodo.IsDone = true

	savedTodo, savedErr := repo.Save(createdTodo)
	if savedErr != nil {
		t.Errorf("エラー: %v", savedErr)
	}

	foundTodo, foundErr := repo.FindByID(strconv.Itoa(createdTodo.TodoID), savedTodo.UserID)
	if foundErr != nil {
		t.Fatalf("エラー: %v", foundErr)
	}
	if foundTodo.IsDone != createdTodo.IsDone {
		t.Errorf("完了状態の切り替えに失敗")
	}
}

func TestDelete_Success(t *testing.T) {

	db.Init()
	repo := TodoRepository{}
	todo := models.Todo{
		Task:   "test",
		UserID: 0,
	}

	createdTodo, createdErr := repo.Create(todo)
	if createdErr != nil {
		t.Fatalf("エラー: %v", createdErr)
	}

	deletedErr := repo.Delete(createdTodo)
	if deletedErr != nil {
		t.Errorf("エラー: %v", deletedErr)
	}

	_, foundErr := repo.FindByID(strconv.Itoa(createdTodo.TodoID), createdTodo.UserID)
	if foundErr == nil {
		t.Errorf("削除に失敗")
	}
}
