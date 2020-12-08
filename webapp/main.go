package main

import (
	"notification_service_webapp/api"
	"notification_service_webapp/util"
	"os"
	"time"

	"notification_service_webapp/env"

	_ "notification_service_webapp/docs" // docs is generated by Swag CLI, you have to import it.

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Notification Service
// @version 1.0
// @description Notification Service Blueprint.

// @contact.name API Support
// @contact.email syedmrizwan@outlook.com

func main() {
	// for production or release mode
	if env.Env.BuildEnv == util.PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	logger := util.GetLogger()

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Route Paths
	api.RegisterRoutes(r.Group("/api/v1"))

	err := r.Run(env.Env.ServerHost + ":" + env.Env.ServerPort) // listen and serve on 0.0.0.0:8080 --> 127.0.0.1:8080

	if err != nil {
		logger.Error(err.Error())
		_ = logger.Sync()
		os.Exit(1)
	}
}
