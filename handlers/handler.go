package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
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
		todo, err := service.CreateService(dto)
		if err!=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(todo)
	}
}

func GetAllTodosHandler(service *services.TodoService) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		vars :=mux.Vars(r)
		pageNumberStr := vars["page_no"]
    	entriesPerPageStr := vars["entries_per_page"]
		sort := vars["sort"]

		pageNumber, err := strconv.Atoi(pageNumberStr)
		if err != nil {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}

		entriesPerPage, err := strconv.Atoi(entriesPerPageStr)
		if err != nil {
			http.Error(w, "Invalid entries per page", http.StatusBadRequest)
			return
		}
		if sort != "ASC" && sort != "DESC"{
			http.Error(w, "Invalid sorting type", http.StatusBadRequest)
		}

		todos, err := service.GetAllTodosService(pageNumber, entriesPerPage, sort)
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
		todo, err := service.GetTodoByIDService(id)

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

        var status *string
        if dto.Status != "" {
            status = &dto.Status
        }
		
		var description *string
        if dto.Description != "" {
            description = &dto.Description
        }

        todo, err := service.UpdateTodoService(id, title, description, status)

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
        if err := service.DeleteTodoService(id); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusNoContent)
    }
}

func GetTodosByUserIDHandler(service *services.TodoService) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		vars :=mux.Vars(r)
		userIdStr:= vars["user_id"]

		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			http.Error(w, "Invalid userId format, userId shoud be an integer", http.StatusBadRequest)
			return
		}

		todos, err := service.GetTodosByUserIDService(userId)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(todos)
	}
}
