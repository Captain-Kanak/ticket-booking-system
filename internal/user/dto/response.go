package dto

import (
	"github.com/google/uuid"
)

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Age   uint8     `json:"age"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
	User  UserResponse
}
