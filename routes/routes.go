package routes

import (
	"todo-backend/handlers"
	"todo-backend/repositories"
	"todo-backend/services"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func RegisterTodoRoutes(router *mux.Router, session *gocql.Session) {
    repo := repositories.NewTodoRepository(session)
    service := services.NewTodoService(repo)

    router.HandleFunc("/todos/create", handlers.CreateTodoHandler(service)).Methods("POST")
    router.HandleFunc("/todos/get", handlers.GetAllTodosHandler(service)).Methods("GET")
    router.HandleFunc("/todos/get/{id}", handlers.GetTodoByIdhandler(service)).Methods("GET")
    router.HandleFunc("/todos/{id}", handlers.UpdateTodoHandler(service)).Methods("PUT")
    router.HandleFunc("/todos/{id}", handlers.DeleteTodoHandler(service)).Methods("DELETE")
}
