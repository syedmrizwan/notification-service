package main

import (
	"log"
	"notification_service_webapp/database"
	"notification_service_webapp/handler"
	"notification_service_webapp/repository"
	"notification_service_webapp/service"

	"github.com/gin-gonic/gin"
)

// will initialize a handler starting from data sources
// which inject into repository layer
// which inject into service layer
// which inject into handler layer
func Inject() (*gin.Engine, error) {
	log.Println("Injecting data sources")

	/*
	 * repository layer
	 */
	db := database.GetConnection()
	notificationRepository := repository.NewNotificationRepository(db)

	/*
	 * service layer
	 */
	notificationService := service.NewNotificationService(&service.NSConfig{
		NotificationRepository: notificationRepository,
	})

	// initialize gin.Engine
	router := gin.New()

	handler.NewHandler(&handler.Config{
		R:                   router,
		NotificationService: notificationService,
		BaseURL:             "/api/v1/",
	})

	return router, nil
}
