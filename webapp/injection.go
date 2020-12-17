package main

import (
	"log"
	"notification_service_webapp/database"
	"notification_service_webapp/env"
	"notification_service_webapp/handler"
	"notification_service_webapp/repository"
	"notification_service_webapp/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
)

// Inject will initialize a handler starting from data sources
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

	conn, err := stan.Connect(
		env.Env.NatsCluster,
		env.Env.NatsClient,
		stan.NatsURL(env.Env.NatsAddress),
	)
	if err != nil {
		os.Exit(1)
	}
	messagingRepository := repository.NewMessagingRepository(conn)

	/*
	 * service layer
	 */
	notificationService := service.NewNotificationService(&service.NSConfig{
		NotificationRepository: notificationRepository,
	})
	messagingService := service.NewMessagingService(&service.MSConfig{
		MessagingRepository: messagingRepository,
	})

	// initialize gin.Engine
	router := gin.New()

	handler.NewHandler(&handler.Config{
		R:                   router,
		NotificationService: notificationService,
		MessagingService:    messagingService,
		BaseURL:             "/api/v1/",
	})

	return router, nil
}
