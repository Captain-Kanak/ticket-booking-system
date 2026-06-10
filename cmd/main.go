package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"time"
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

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" validate:"required" gorm:"type:varchar(100);not null"`
	Email     string         `json:"email" validate:"required,email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string         `json:"password" validate:"required,min=6" gorm:"type:varchar(255);not null"`
	Age       uint8          `json:"age"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=ticket_booking port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	fmt.Println("Database connected successfully!")

	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.Validator = &CustomValidator{validator: validator.New()}

	// * api routes
	e.GET("/", func(c *echo.Context) error {
		// return c.String(http.StatusOK, "Hello, World!")

		return c.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Ticket Booking System - Server is running successfully!",
		})
	})

	e.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Server is healthy!",
		})
	})

	e.POST("/users", func(c *echo.Context) (err error) {
		newUser := new(User)

		// * bind and validate
		if err = c.Bind(newUser); err != nil {
			return c.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Invalid request body!",
			})
		}

		if err = c.Validate(newUser); err != nil {
			fmt.Println(err.Error())

			return c.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Validation failed!",
			})
		}

		// * save user information to database
		tx := db.Create(newUser)

		if tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Message: "Failed to create user!",
			})
		}

		if tx.RowsAffected == 0 {
			return c.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Message: "Failed to create user!",
			})
		}

		return c.JSON(http.StatusOK, Response{
			Success: true,
			Message: "User created successfully!",
			Data:    newUser,
		})
	})

	// * start server
	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
