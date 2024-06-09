package models

import "github.com/gocql/gocql"

type Todo struct {
	ID 		gocql.UUID `json:"id"`
	Title	string     `json:"title"`
	Status	string	   `json:"status"`
}