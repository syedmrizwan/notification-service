package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"notification_service_webapp/model"
	"time"

	"github.com/gin-gonic/gin"
)

// GetNotifications godoc
// @Tags Notifications
// @Summary Get All Notifications
// @Accept json
// @Produce json
// @Success 200 {object} []model.Notification
// @Router /api/v1/notifications [GET]
func (h *Handler) GetNotifications(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, time.Minute)
	defer cancel()
	notifications := make([]model.Notification, 0)
	notifications, err := h.NotificationService.GetAll(ctx)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}

// CreateNotifications godoc
// @Tags Notifications
// @Summary Create Notifications
// @Accept json
// @Produce json
// @Param payload body model.NotificationPostBody true "description"
// @Success 200 {object} model.Notification
// @Router /api/v1/notifications [POST]
func (h *Handler) CreateNotifications(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, time.Minute)
	defer cancel()

	var notificationPostBody model.NotificationPostBody
	if err := c.ShouldBindJSON(&notificationPostBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	notification, err := h.NotificationService.Create(ctx, &notificationPostBody)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	message, _ := json.Marshal(notification)
	if err := h.MessagingService.Write(notification.Priority, message); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"notification": notification})
}

// CreateBulkNotifications godoc
// @Tags Notifications
// @Summary Create Bulk Notifications
// @Accept json
// @Produce json
// @Param payload body model.BulkNotificationBody true "description"
// @Success 200 {object} []model.Notification
// @Router /api/v1/bulk-notifications [POST]
func (h *Handler) CreateBulkNotifications(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c, 2*time.Minute)
	defer cancel()

	var bulkNotificationBody model.BulkNotificationBody
	if err := c.ShouldBindJSON(&bulkNotificationBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	notifications, err := h.NotificationService.BulkCreate(ctx, &bulkNotificationBody)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for _, notification := range notifications {
		message, _ := json.Marshal(notification)
		if err := h.MessagingService.Write(notification.Priority, message); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{"notifications": notifications})
}
