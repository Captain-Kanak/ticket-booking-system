package event

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	eventRepo := NewRepository(db)
	eventService := NewService(eventRepo)
	eventHandler := NewHandler(eventService)

	api := e.Group("/api/v1/events")

	api.POST("", eventHandler.Create)
	api.GET("", eventHandler.GetAll)
	api.GET("/:id", eventHandler.GetByID)
}
