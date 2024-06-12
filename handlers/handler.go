package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-backend/dtos"
	"todo-backend/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoResponse struct {
    Message string      `json:"message"`
    Todo    interface{} `json:"todo"`
}

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
		response := TodoResponse{
            Message: "Todo created successfully",
            Todo:    todo,
        }
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
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
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todos)
	}
}

func GetTodoByIdhandler(service *services.TodoService) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		vars :=mux.Vars(r)
		idHex := vars["id"]
		id, err := primitive.ObjectIDFromHex(idHex)
		if err!=nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		todo, err := service.GetTodoByIDService(id)

		fmt.Printf( "%+v\n",todo)

		if err!=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := TodoResponse{
            Message: "Todos retrieved successfully",
            Todo:    todo,
        }
		json.NewEncoder(w).Encode(response)
	}
}

func UpdateTodoHandler(service *services.TodoService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        idHex := vars["id"]
		id, err := primitive.ObjectIDFromHex(idHex)
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
		response := TodoResponse{
            Message: "Todo updated successfully",
            Todo:    todo,
        }
        w.WriteHeader(http.StatusNoContent)
        json.NewEncoder(w).Encode(response)
    }
}

func DeleteTodoHandler(service *services.TodoService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        idHex := vars["id"]
		id, err := primitive.ObjectIDFromHex(idHex)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        if err := service.DeleteTodoService(id); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
		response := TodoResponse{
            Message: "Todo deleted successfully",
        }
        w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(response)
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
		response := TodoResponse{
            Message: "Todos retrieved successfully",
			Todo: todos,
        }
		json.NewEncoder(w).Encode(response)
	}
}
