package main

import (
	"ticket-booking-system/internal/config"
	"ticket-booking-system/internal/server"

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

	// * start server
	server.StartServer(e, cfg, db)
}
