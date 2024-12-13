package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port int64
}

// Envs is the global configuration for the application.
var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		Port: getEnvInt("PORT", 3000),
	}
}

func getEnvInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
