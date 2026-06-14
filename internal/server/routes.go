package server

import (
	"ticket-booking-system/internal/event"
	"ticket-booking-system/internal/user"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RoutesHandler(e *echo.Echo, db *gorm.DB) {
	// * user routes
	db.AutoMigrate(user.User{})
	user.Routes(e, db)

	// * event routes
	db.AutoMigrate(event.Event{})
	event.Routes(e, db)
}
