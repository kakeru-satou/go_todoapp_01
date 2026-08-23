package models

import "errors"

type Todo struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}

type CreateTodoRequest struct {
	Task string `json:"task"`
}

type PatchTodoRequest struct {
	Task   *string `json:"task"`
	IsDone *bool   `json:"isDone"`
}

var ErrTaskRequired = errors.New("task is required")
