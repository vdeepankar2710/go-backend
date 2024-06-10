package services

import (
	"time"
	"todo-backend/dtos"
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

func (s *TodoService) CreateService (dto dtos.CreateTodoDTO) (*models.Todo, error){
	todo := &models.Todo {
		ID:          gocql.TimeUUID(),
		UserID:      dto.UserID,
		Title:       dto.Title,
		Description: dto.Description,
		Status:      dto.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err:= s.repository.CreateRepo(todo); err!=nil{
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) GetAllTodosService (pageNumber int, entriesPerPage int) ([]models.Todo, error){
	todos, err := s.repository.GetAllRepo(pageNumber, entriesPerPage)
	if err != nil{
		return nil, err
	}
	return todos, nil
}

func (s *TodoService) GetTodoByIDService (id gocql.UUID) (*models.Todo, error){
	todo, err := s.repository.GetTodoByIDRepo(id)
	if err != nil{
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) UpdateTodoService (id gocql.UUID, title *string, description *string, status *string) (*models.Todo, error){
	todo, err := s.repository.GetTodoByIDRepo(id)
	if err!=nil{
		return nil, err
	}

	if title != nil {
        todo.Title = *title
    }
    if status != nil {
        todo.Status = *status
    }

	if description!=nil{
		todo.Description = *description
	}

	todo.UpdatedAt = time.Now()

	if err:= s.repository.UpdateTodoRepo(todo); err!=nil{
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) DeleteTodoService (id gocql.UUID) error{
	if err := s.repository.DeleteTodoRepo(id); err!=nil{
		return err
	}
	return nil
}


func (s *TodoService) GetTodosByUserIDService (userId int) ([]models.Todo, error){
	todos, err := s.repository.GetTodoByUserIdRepo(userId)
	if err!=nil{
		return nil, err
	}
	return todos, err
	
}