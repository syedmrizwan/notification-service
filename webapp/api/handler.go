package api

import (
	"context"
	"net/http"
	"notification_service_webapp/database"
	"notification_service_webapp/model"
	"notification_service_webapp/util"

	"time"

	"github.com/gin-gonic/gin"
)

type NotificationPostBody struct {
	NotificationText string `json:"notification_text"`
	Priority         string `json:"Priority"`
	UserId           int    `json:"user_id"`
	NotificationMode string `json:"notification_mode"`
}

type ResponseBody struct {
	Message string `json:"message"`
}

// getNotifications godoc
// @Tags Notifications
// @Summary Get All Notifications
// @Accept json
// @Produce json
// @Success 200 {object} []model.Notification
// @Router /api/v1/notifications [GET]
func getNotifications(c *gin.Context) {
	logger := util.GetLogger()
	defer logger.Sync()
	logger.Info("Message received for getNotifications")
	ctx, cancel := context.WithTimeout(c, time.Minute)
	defer cancel()
	db := database.GetConnection()

	notifications := make([]model.Notification, 0)

	if err := db.ModelContext(ctx, &notifications).Select(); err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// createNotifications godoc
// @Tags Notifications
// @Summary Create Notifications
// @Accept json
// @Produce json
// @Param payload body NotificationPostBody true "description"
// @Success 200 {object} model.Notification
// @Router /api/v1/notifications [POST]
func createNotifications(c *gin.Context) {
	logger := util.GetLogger()
	defer logger.Sync()
	logger.Info("Message received for createNotifications")
	ctx, cancel := context.WithTimeout(c, time.Minute)
	defer cancel()
	db := database.GetConnection()

	var notificationPostBody NotificationPostBody
	if err := c.ShouldBindJSON(&notificationPostBody); err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	notificationHandler := model.NotificationHandler{Name: notificationPostBody.NotificationMode}
	if err := db.ModelContext(ctx, &notificationHandler).Where("name = ?name").Select(); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusNotFound, ResponseBody{Message: "Notification Mode does not exists"})
		return
	}

	notificationText := model.NotificationText{Message: notificationPostBody.NotificationText}
	if _, err := db.ModelContext(ctx, &notificationText).Where("message = ?message").SelectOrInsert(); err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	notification := model.Notification{
		Priority:              notificationPostBody.Priority,
		Status:                util.SENDING,
		UserId:                notificationPostBody.UserId,
		NotificationHandlerID: notificationHandler.ID,
		NotificationTextID:    notificationText.ID,
	}

	if err := db.Insert(&notification); err != nil {
		logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, notification)
}
