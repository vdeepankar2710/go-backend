package main

import (
	"fmt"
	"net/http"
	"todo-backend/config"
	"todo-backend/routes"

	"github.com/gorilla/mux"
)

func main(){
	
	session, err := config.NewSession()
	if err!=nil{
		fmt.Printf("Failed to connect to DB %v", err)
	}
	defer session.Close()

	router := mux.NewRouter()

    routes.RegisterTodoRoutes(router, session)

    http.ListenAndServe(":8080", router)
	
	fmt.Printf("Failed to connect to DB %v", err)
	
}