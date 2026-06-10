package main

import (
	"fmt"
	"net/http"
	"ticket-booking-system/internal/httpresponse"
	"ticket-booking-system/internal/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/driver/postgres"
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

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=ticket_booking port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	// * database connection
	func() {
		if err != nil {
			panic("failed to connect database")
		}

		db.AutoMigrate(user.User{})

		fmt.Println("Database connected successfully!")
	}()

	e := echo.New()

	// * echo middleware & validator
	func() {
		e.Use(middleware.RequestLogger())

		e.Validator = &CustomValidator{validator: validator.New()}
	}()

	// * basics api routes
	func() {
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
	}()

	// * user routes
	func() {
		userRepo := user.NewRepository(db)
		userService := user.NewService(userRepo)
		userHandler := user.NewHandler(userService)

		e.POST("/users", userHandler.CreateUser)
	}()

	// * start server
	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
