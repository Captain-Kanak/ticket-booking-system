package event

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Venue struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string         `json:"title" gorm:"type:varchar(255);not null"`
	Description string         `json:"description" gorm:"type:text"`
	City        string         `json:"city" gorm:"type:varchar(255);not null"`
	Capacity    int            `json:"capacity" gorm:"type:int;not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type Event struct {
	ID          uint
	Title       string
	Description string
	Category    string
	VenueID     uint
	Venue       Venue
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
