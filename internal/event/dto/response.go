package dto

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	ID               uuid.UUID `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Location         string    `json:"location"`
	StartDate        time.Time `json:"start_date"`
	TotalTickets     int       `json:"total_tickets"`
	AvailableTickets int       `json:"available_tickets"`
	Price            float64   `json:"price"`
	CreatedAt        time.Time `json:"created_at"`
}
