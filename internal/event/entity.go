package event

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID               uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title            string         `json:"title" gorm:"type:varchar(255);not null"`
	Description      string         `json:"description" gorm:"type:text;not null"`
	Location         string         `json:"location" gorm:"type:text;not null"`
	StartDate        time.Time      `json:"start_date" gorm:"type:timestamp;not null"`
	TotalTickets     int            `json:"total_tickets" gorm:"type:int;not null"`
	AvailableTickets int            `json:"available_tickets" gorm:"type:int;not null"`
	Price            float64        `json:"price" gorm:"type:float;not null"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}

// type Venue struct {
// 	ID          uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
// 	Title       string         `json:"title" gorm:"type:varchar(255);not null"`
// 	Description string         `json:"description" gorm:"type:text"`
// 	City        string         `json:"city" gorm:"type:varchar(255);not null"`
// 	Capacity    int            `json:"capacity" gorm:"type:int;not null"`
// 	CreatedAt   time.Time      `json:"created_at"`
// 	UpdatedAt   time.Time      `json:"updated_at"`
// 	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
// }
