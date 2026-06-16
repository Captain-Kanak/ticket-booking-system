package event

import (
	"errors"
	"fmt"
	"net/http"
	"ticket-booking-system/internal/event/dto"
	"ticket-booking-system/internal/httpresponse"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
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

func (h *handler) GetByID(c *echo.Context) (err error) {
	id := c.Param("id")

	parsedId, err := uuid.Parse(id)

	if err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, httpresponse.Response{
			Success: false,
			Message: "Invalid event ID",
			Error:   err.Error(),
		})
	}

	res, err := h.service.GetByID(parsedId)

	if err != nil {
		fmt.Println(err.Error())

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, httpresponse.Response{
				Success: false,
				Message: "Event not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, httpresponse.Response{
			Success: false,
			Message: "Failed to get event",
		})
	}

	return c.JSON(http.StatusOK, httpresponse.Response{
		Success: true,
		Message: "Event fetched successfully",
		Data:    res,
	})
}
