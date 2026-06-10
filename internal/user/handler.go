package user

import (
	"fmt"
	"github.com/labstack/echo/v5"
	"net/http"
	"ticket-booking-system/internal/httpresponse"
	"ticket-booking-system/internal/user/dto"
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
			Message: "Failed to bind request body",
		})
	}

	if err = c.Validate(req); err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, httpresponse.Response{
			Success: false,
			Message: "Validation failed!",
		})
	}

	res, err := h.service.CreateUser(req)

	if err != nil {
		fmt.Println(err.Error())

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
