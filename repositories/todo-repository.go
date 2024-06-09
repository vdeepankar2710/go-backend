package repositories

import (
	"todo-backend/models"

	"github.com/gocql/gocql"
)

type TodoRepository struct{
	session *gocql.Session
}

func NewTodoRepository(session *gocql.Session) *TodoRepository {
	return &TodoRepository{session: session}
}

func (r *TodoRepository) Create(todo *models.Todo) error{
	return r.session.Query("INSERT INTO todos (id, title, status) VALUES (?, ?, ?)", todo.ID, todo.Title, todo.Status).Exec()
}

func (r *TodoRepository) GetAll() ([]models.Todo, error){
	var todosArr []models.Todo

	iter:= r.session.Query("SELECT id, title, status FROM todos").Iter()
	var todo models.Todo
	for iter.Scan(&todo.ID, &todo.Title, &todo.Status){
		todosArr = append(todosArr, todo)
	}

	if err := iter.Close(); err!=nil{
		return nil, err
	}
	return todosArr, nil
}

func (r *TodoRepository) GetTodoByID(id gocql.UUID) (*models.Todo, error) {
	var todo models.Todo

	if err := r.session.Query("SELECT id, title, status FROM todos where id = ?", id).Scan(&todo.ID, &todo.Title, &todo.Status); err!=nil{
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) UpdateTodo(todo *models.Todo) error {
	return r.session.Query("UPDATE todos SET title = ?, status = ? WHERE id = ?", todo.Title, todo.Status, todo.ID).Exec()
}

func (r *TodoRepository) DeleteTodo(id gocql.UUID) error {
	return r.session.Query("DELETE FROM todos WHERE id = ?", id).Exec()
}