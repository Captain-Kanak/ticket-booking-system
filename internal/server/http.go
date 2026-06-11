package server

import (
	"fmt"
	"net/http"
	"ticket-booking-system/internal/config"
	"ticket-booking-system/internal/httpresponse"
	"ticket-booking-system/internal/user"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func StartServer(e *echo.Echo, cfg *config.EnvConfig, db *gorm.DB) {
	// * basics api routes
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

	port := fmt.Sprintf(":%s", cfg.Port)

	if err := e.Start(port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
