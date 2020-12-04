package util

import (
	"notification_service_webapp/env"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {

	switch env.Env.BuildEnv {
	case DEVELOPMENT:
		logger, _ = zap.NewDevelopment()
		return
	case PRODUCTION:
		// first, we create a configuration (i.e. a builder) for our logger
		config := zap.NewProductionConfig()

		// we configure the destination for our log, in this case, a file
		config.OutputPaths = []string{"./logs.log"}

		// we build the logger
		logger, _ = config.Build()
		return
	}
}

func GetLogger() *zap.Logger {
	return logger
}
