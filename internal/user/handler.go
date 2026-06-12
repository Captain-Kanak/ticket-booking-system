package user

import (
	"errors"
	"fmt"
	"net/http"
	"ticket-booking-system/internal/httpresponse"
	"ticket-booking-system/internal/user/dto"

	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateUser(c *echo.Context) (err error) {
	var req = new(dto.CreateRequest)

	if err = c.Bind(req); err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, httpresponse.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	if err = c.Validate(req); err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, httpresponse.Response{
			Success: false,
			Message: "Validation failed!",
			Error:   err.Error(),
		})
	}

	res, err := h.service.CreateUser(req)

	if err != nil {
		fmt.Println(err.Error())

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusConflict, httpresponse.Response{
				Success: false,
				Message: "User already exists with this email",
			})
		}

		return c.JSON(http.StatusInternalServerError, httpresponse.Response{
			Success: false,
			Message: "Failed to create user",
		})
	}

	return c.JSON(http.StatusCreated, httpresponse.Response{
		Success: true,
		Message: "User created successfully",
		Data:    res,
	})
}

func (h *handler) LoginUser(c *echo.Context) (err error) {
	var req = new(dto.LoginRequest)

	if err = c.Bind(req); err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, httpresponse.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	if err = c.Validate(req); err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, httpresponse.Response{
			Success: false,
			Message: "Validation failed!",
			Error:   err.Error(),
		})
	}

	res, err := h.service.LoginUser(req)

	if err != nil {
		fmt.Println(err.Error())

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, httpresponse.Response{
				Success: false,
				Message: "User not found with this email",
			})
		}

		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return c.JSON(http.StatusUnauthorized, httpresponse.Response{
				Success: false,
				Message: "Invalid credentials",
			})
		}

		return c.JSON(http.StatusInternalServerError, httpresponse.Response{
			Success: false,
			Message: "Failed to login user",
		})
	}

	return c.JSON(http.StatusOK, httpresponse.Response{
		Success: true,
		Message: "User logged in successfully",
		Data:    res,
	})
}
