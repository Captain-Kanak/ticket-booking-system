package event

import (
	"fmt"
	"net/http"
	"ticket-booking-system/internal/event/dto"
	"ticket-booking-system/internal/httpresponse"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{service}
}

func (h *handler) Create(c *echo.Context) (err error) {
	var req dto.CreateRequest

	if err = c.Bind(&req); err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, httpresponse.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	if err = c.Validate(&req); err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, httpresponse.Response{
			Success: false,
			Message: "Validation failed!",
			Error:   err.Error(),
		})
	}

	res, err := h.service.Create(&req)

	if err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusInternalServerError, httpresponse.Response{
			Success: false,
			Message: "Failed to create event",
		})
	}

	return c.JSON(http.StatusOK, httpresponse.Response{
		Success: true,
		Message: "Event created successfully",
		Data:    res,
	})
}

func (h *handler) GetAll(c *echo.Context) (err error) {
	res, err := h.service.GetAll()

	if err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusInternalServerError, httpresponse.Response{
			Success: false,
			Message: "Failed to get events",
		})
	}

	return c.JSON(http.StatusOK, httpresponse.Response{
		Success: true,
		Message: "Events fetched successfully",
		Data:    res,
	})
}
