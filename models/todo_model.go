package models

type Todo struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}

type CreateTodoRequest struct {
	Task string `json:"task"`
}
