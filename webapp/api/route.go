package api

import "github.com/gin-gonic/gin"

//RegisterRoutes register api routes
func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/notifications", getNotifications)
	router.POST("/notifications", createNotifications)
	router.POST("/bulk-notifications", createBulkNotifications)
}
