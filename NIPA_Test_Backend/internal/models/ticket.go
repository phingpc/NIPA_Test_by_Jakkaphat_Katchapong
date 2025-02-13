package models

import "time"

type Ticket struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Contact     string    `json:"contact"`
	CreatedAt   time.Time `json:"created"`
	UpdatedAt   time.Time `json:"updated"`
	Status      string    `json:"status"`
}

type TicketInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Contact     string `json:"contact"`
}
