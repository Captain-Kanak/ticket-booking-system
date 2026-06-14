package dto

import "time"

type CreateRequest struct {
	Title        string    `json:"title" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	Location     string    `json:"location" validate:"required"`
	StartDate    time.Time `json:"start_date" validate:"required"`
	TotalTickets int       `json:"total_tickets" validate:"required"`
	Price        float64   `json:"price" validate:"required"`
}

type UpdateRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Location    string    `json:"location" validate:"required"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	Price       float64   `json:"price" validate:"required"`
}
