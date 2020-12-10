package handler

import (
	"notification_service_webapp/model"

	"github.com/gin-gonic/gin"
)

// Handler struct holds required services for handler to function
type Handler struct {
	NotificationService model.NotificationService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R                   *gin.Engine
	NotificationService model.NotificationService
	BaseURL             string
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		NotificationService: c.NotificationService,
	} // currently has no properties

	// Create an account group
	g := c.R.Group(c.BaseURL)

	g.GET("/notifications", h.GetNotifications)
	g.POST("/notifications", h.CreateNotifications)
	g.POST("/bulk-notifications", h.CreateBulkNotifications)

}
