package dtos

type CreateTodoDTO struct {
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateTodoDTO struct {
	Title       string `json:"title"`
	Status      string `json:"status"`
	Description string `json:"description"`
}