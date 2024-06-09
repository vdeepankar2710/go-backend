package handlers

import (
	"encoding/json"
	"net/http"
	"todo-backend/dtos"
	"todo-backend/services"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func CreateTodoHandler(service *services.TodoService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var dto dtos.CreateTodoDTO
		if err:= json.NewDecoder(r.Body).Decode(&dto); err!=nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		todo, err := service.Create(dto.Title)
		if err!=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(todo)
	}
}

func GetAllTodosHandler(service *services.TodoService) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		todos, err := service.GetAllTodos()
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(todos)
	}
}

func GetTodoByIdhandler(service *services.TodoService) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		vars :=mux.Vars(r)
		id, err := gocql.ParseUUID(vars["id"])
		if err!=nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		todo, err := service.GetTodoByID(id)

		if err!=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(todo)
	}
}

func UpdateTodoHandler(service *services.TodoService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, err := gocql.ParseUUID(vars["id"])
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        var dto dtos.UpdateTodoDTO
        if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
		var title *string
        if dto.Title != "" {
            title = &dto.Title
        }

        var completed *string
        if dto.Status != "" {
            completed = &dto.Status
        }

        todo, err := service.UpdateTodo(id, title, completed)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(todo)
    }
}

func DeleteTodoHandler(service *services.TodoService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, err := gocql.ParseUUID(vars["id"])
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        if err := service.DeleteTodo(id); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusNoContent)
    }
}
