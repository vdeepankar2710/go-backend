package dtos

type CreateTodoDTO struct {
	Title string `json:"title"`
}

type UpdateTodoDTO struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}