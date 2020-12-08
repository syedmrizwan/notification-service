package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type envFile struct {
	DbName        string
	DbUsername    string
	DbPassword    string
	DbHost        string
	DbPort        string
	BuildEnv      string
	ServerHost    string
	ServerPort    string
	DbPoolSize    int
	NatsAddress   string
	NatsCluster   string
	NatsClient    string
	IsContainer   bool
	RedisPassword string
	RedisAddress  string
}

func (e *envFile) GetAddr() string {
	return e.DbHost + ":" + e.DbPort
}

var Env *envFile

func init() {
	_ = godotenv.Load()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	IsContainer, _ := strconv.ParseBool(os.Getenv("IS_CONTAINER"))
	DbPoolSize, _ := strconv.Atoi(os.Getenv("DB_POOL_SIZE"))

	Env = &envFile{
		DbName:        os.Getenv("DB_NAME"),
		DbUsername:    os.Getenv("DB_USERNAME"),
		DbPassword:    os.Getenv("DB_PASSWORD"),
		DbHost:        os.Getenv("DB_HOST"),
		DbPort:        os.Getenv("DB_PORT"),
		BuildEnv:      os.Getenv("BUILD_ENV"),
		ServerHost:    os.Getenv("SERVER_HOST"),
		ServerPort:    os.Getenv("SERVER_PORT"),
		NatsAddress:   os.Getenv("NATS_URL"),
		NatsCluster:   os.Getenv("NATS_CLUSTER_ID"),
		NatsClient:    os.Getenv("NATS_CLIENT_ID"),
		DbPoolSize:    DbPoolSize,
		IsContainer:   IsContainer,
		RedisAddress:  os.Getenv("REDIS_ADDRESS"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}
}
