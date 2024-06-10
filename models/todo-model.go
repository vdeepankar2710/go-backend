package models

import (
	"time"

	"github.com/gocql/gocql"
)

type Todo struct {
	ID          gocql.UUID `json:"id"`
	UserID      int 	   `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}