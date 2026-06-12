package user

import (
	"ticket-booking-system/internal/auth"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	userRepo := NewRepository(db)
	jwtService := auth.NewJWTService("")
	userService := NewService(userRepo, jwtService)
	userHandler := NewHandler(userService)

	api := e.Group("/api/v1")

	api.POST("/register", userHandler.CreateUser)
	api.POST("/login", userHandler.LoginUser)
}
