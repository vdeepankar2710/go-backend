package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-backend/config"
	"todo-backend/routes"

	"github.com/gorilla/mux"
)

func main(){

	MONGO_URI := "mongodb://localhost:27017"
	db, err := config.ConnectDB(MONGO_URI)
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

	router := mux.NewRouter()

    routes.RegisterTodoRoutes(router, db)

    log.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", router)
	
	fmt.Printf("Failed to connect to DB %v", err)
	
}