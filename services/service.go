package services

import (
	"time"
	"todo-backend/dtos"
	"todo-backend/models"
	"todo-backend/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoService struct {
	repository *repositories.TodoRepository
}

func NewTodoService(repository *repositories.TodoRepository) *TodoService {
	return &TodoService{repository: repository}
}

func (s *TodoService) CreateService (dto dtos.CreateTodoDTO) (*models.Todo, error){
	todo := &models.Todo {
		ID:          primitive.NewObjectID(),
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

func (s *TodoService) GetAllTodosService (pageNumber int, entriesPerPage int, sort string) ([]models.Todo, error){
	todos, err := s.repository.GetAllRepo(pageNumber, entriesPerPage, sort)
	if err != nil{
		return nil, err
	}
	return todos, nil
}

func (s *TodoService) GetTodoByIDService (id primitive.ObjectID) (*models.Todo, error){
	todo, err := s.repository.GetTodoByIDRepo(id)
	if err != nil{
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) UpdateTodoService (id primitive.ObjectID, title *string, description *string, status *string) (*models.Todo, error){
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

func (s *TodoService) DeleteTodoService (id primitive.ObjectID) error{
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