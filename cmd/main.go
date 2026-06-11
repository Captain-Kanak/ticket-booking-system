package main

import (
	"fmt"
	"net/http"
	"ticket-booking-system/internal/config"
	"ticket-booking-system/internal/httpresponse"
	"ticket-booking-system/internal/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}

	return nil
}

func main() {
	cfg := config.LoadEnv()

	db := config.ConnectDatabase(cfg)

	e := echo.New()

	// * echo middleware & validator
	e.Use(middleware.RequestLogger())

	e.Validator = &CustomValidator{validator: validator.New()}

	// * basics api routes
	e.GET("/", func(c *echo.Context) error {
		// return c.String(http.StatusOK, "Hello, World!")

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

	// * start server
	port := fmt.Sprintf(":%s", cfg.Port)

	if err := e.Start(port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
