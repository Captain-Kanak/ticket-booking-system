package venue

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Venue struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string         `json:"name" gorm:"type:varchar(255);not null"`
	Address   string         `json:"address" gorm:"type:varchar(255);not null"`
	City      string         `json:"city" gorm:"type:varchar(255);not null"`
	Capacity  int            `json:"capacity" gorm:"type:int;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
