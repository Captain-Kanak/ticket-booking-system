package server

import (
	"net/http"
	"ticket-booking-system/internal/event"
	"ticket-booking-system/internal/httpresponse"
	"ticket-booking-system/internal/user"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RoutesHandler(e *echo.Echo, db *gorm.DB) {
	// * db migrations
	db.AutoMigrate(user.User{}, event.Event{})

	// * basics routes
	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, httpresponse.Response{
			Success: true,
			Message: "Ticket Booking System - Server is running successfully!",
		})
	})

	e.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, httpresponse.Response{
			Success: true,
			Message: "Server is healthy!",
		})
	})

	// * user routes
	user.Routes(e, db)

	// * event routes
	event.Routes(e, db)
}
