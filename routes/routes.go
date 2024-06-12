package routes

import (
	"todo-backend/handlers"
	"todo-backend/repositories"
	"todo-backend/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterTodoRoutes(router *mux.Router, db *mongo.Database) {
    repo := repositories.NewTodoRepository(db)
    service := services.NewTodoService(repo)

    router.HandleFunc("/todos/create", handlers.CreateTodoHandler(service)).Methods("POST")
    router.HandleFunc("/todos/get/{page_no}/{entries_per_page}/{sort}", handlers.GetAllTodosHandler(service)).Methods("GET")
    router.HandleFunc("/todos/get/{id}", handlers.GetTodoByIdhandler(service)).Methods("GET")
    router.HandleFunc("/todos/update/{id}", handlers.UpdateTodoHandler(service)).Methods("PUT")
    router.HandleFunc("/todos/delete/{id}", handlers.DeleteTodoHandler(service)).Methods("DELETE")
}
