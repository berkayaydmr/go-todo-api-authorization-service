package common

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

type Environment struct {
	RouterUrl string
	RedisUrl  string
	Debug     bool
	SecretKey string
	Network   string
}

func GetEnvironment() *Environment {
	err := godotenv.Load(".env")
	if err != nil {
		zap.S().Error("Error: ", err)
		return nil
	}

	appHost := os.Getenv("APPLICATION_HOST")
	appPort := os.Getenv("APPLICATION_PORT")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	network := os.Getenv("NETWORK")
	routerUrl := appHost + ":" + appPort
	redisUrl := redisHost + ":" + redisPort

	secretKey := os.Getenv("ACCESS_KEY")

	var debug bool
	if os.Getenv("DEBUG") == "true" {
		debug = true
	}

	return &Environment{
		RouterUrl: routerUrl,
		RedisUrl:  redisUrl,
		Debug:     debug,
		SecretKey: secretKey,
		Network:   network,
	}
}
