package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"notification_service_webapp/model"
	"notification_service_webapp/util"
	"time"

	"github.com/gin-gonic/gin"
)

var errNotificationMode = errors.New("notification Mode does not exists")

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

	if !util.ContainsString(util.SupportedPriority, notificationPostBody.Priority) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Priority can either be High, Medium or Low"})
		return
	}

	notification, err := h.NotificationService.Create(ctx, &notificationPostBody)
	if err != nil {
		switch err.Error() {
		case errNotificationMode.Error():
			c.JSON(http.StatusNotFound, gin.H{"message": "Notification Mode does not exists"})
		default:
			c.AbortWithError(http.StatusInternalServerError, err)
		}
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
		switch err.Error() {
		case errNotificationMode.Error():
			c.JSON(http.StatusNotFound, gin.H{"message": "Notification Mode does not exists"})
		default:
			c.AbortWithError(http.StatusInternalServerError, err)
		}
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
