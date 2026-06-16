package server

import (
	"fmt"
	"ticket-booking-system/internal/config"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
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

func Start(cfg *config.EnvConfig, db *gorm.DB) {
	e := echo.New()

	// * echo middleware & validator
	e.Use(middleware.RequestLogger())
	e.Validator = &CustomValidator{validator: validator.New()}

	// * db migrations & routes handler
	RoutesHandler(e, db)

	port := fmt.Sprintf(":%s", cfg.Port)

	if err := e.Start(port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
