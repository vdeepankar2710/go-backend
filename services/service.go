package services

import (
	"todo-backend/models"
	"todo-backend/repositories"

	"github.com/gocql/gocql"
)

type TodoService struct {
	repository *repositories.TodoRepository
}

func NewTodoService(repository *repositories.TodoRepository) *TodoService {
	return &TodoService{repository: repository}
}

func (s *TodoService) Create(title string) (*models.Todo, error){
	todo := &models.Todo{
		ID:gocql.TimeUUID(),
		Title: title,
		Status: "Not Competed",
	}
	if err:= s.repository.Create(todo); err!=nil{
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) GetAllTodos () ([]models.Todo, error){
	todos, err := s.repository.GetAll()
	if err != nil{
		return nil, err
	}

	return todos, nil
}

func (s *TodoService) GetTodoByID (id gocql.UUID) (*models.Todo, error){
	todo, err := s.repository.GetTodoByID(id)
	if err != nil{
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) UpdateTodo (id gocql.UUID, title *string, status *string) (*models.Todo, error){
	todo, err := s.repository.GetTodoByID(id)
	if err!=nil{
		return nil, err
	}

	if title != nil {
        todo.Title = *title
    }
    if status != nil {
        todo.Status = *status
    }
	if err:= s.repository.UpdateTodo(todo); err!=nil{
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) DeleteTodo (id gocql.UUID) error{
	if err := s.repository.DeleteTodo(id); err!=nil{
		return err
	}
	return nil
}