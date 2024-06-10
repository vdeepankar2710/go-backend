package repositories

import (
	"todo-backend/errors"
	"todo-backend/models"

	"github.com/gocql/gocql"
)

type TodoRepository struct{
	session *gocql.Session
}

func NewTodoRepository(session *gocql.Session) *TodoRepository {
	return &TodoRepository{session: session}
}

func (r *TodoRepository) CreateRepo(todo *models.Todo) error{
	return r.session.Query("INSERT INTO todos (id, user_id, title, description, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)", 
							todo.ID, todo.UserID, todo.Title, todo.Description, todo.Status, todo.CreatedAt, todo.UpdatedAt).Exec()
}

func (r *TodoRepository) GetAllRepo(pageNumber int, entriesPerPage int, sort string) ([]models.Todo, error){
	var todosArr []models.Todo
	var count int

	offset := (pageNumber - 1) * entriesPerPage

    err := r.session.Query("SELECT COUNT(*) FROM todos").Scan(&count)
    if err != nil {
        return nil, err
    }

	if offset > count{
		return nil, errors.ErrInvalidOffset
	}

	iter:= r.session.Query("SELECT * FROM todos LIMIT %d OFFSET %d ORDER BY created_at ?", entriesPerPage, offset, sort).Iter()
	var todo models.Todo
	for iter.Scan(
		&todo.ID,
        &todo.UserID,
        &todo.Title,
        &todo.Description,
        &todo.Status,
        &todo.CreatedAt,
        &todo.UpdatedAt){
		todosArr = append(todosArr, todo)
	}

	if err := iter.Close(); err!=nil{
		return nil, err
	}
	return todosArr, nil
}

func (r *TodoRepository) GetTodoByIDRepo(id gocql.UUID) (*models.Todo, error) {
	var todo models.Todo

	if err := r.session.Query("SELECT id, title, status FROM todos where id = ?", id).Scan(&todo.ID, &todo.Title, &todo.Status); err!=nil{
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) UpdateTodoRepo(todo *models.Todo) error {
	return r.session.Query("UPDATE todos SET title = ?, status = ? WHERE id = ?", todo.Title, todo.Status, todo.ID).Exec()
}

func (r *TodoRepository) DeleteTodoRepo(id gocql.UUID) error {
	return r.session.Query("DELETE FROM todos WHERE id = ?", id).Exec()
}

func (r *TodoRepository) GetTodoByUserIdRepo(userId int) ([]models.Todo, error){
	var todosArr []models.Todo

	iter:= r.session.Query("SELECT * FROM todos WHERE user_id = ?", userId).Iter()
	var todo models.Todo
	for iter.Scan(
		&todo.ID,
        &todo.UserID,
        &todo.Title,
        &todo.Description,
        &todo.Status,
        &todo.CreatedAt,
        &todo.UpdatedAt){
		todosArr = append(todosArr, todo)
	}

	if err := iter.Close(); err!=nil{
		return nil, err
	}
	return todosArr, nil
}