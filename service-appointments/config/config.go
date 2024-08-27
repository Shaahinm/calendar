package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	DbName          string
	DbHost          string
	DbPort          string
	DbUser          string
	DbPass          string
	ServerName      string
	DbType          string
	ResponseVersion int
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Port:            getEnv("PORT", "8080"),
		DbName:          getEnv("DB_NAME", "test.db"),
		DbHost:          getEnv("DB_HOST", "127.0.0.1"),
		DbPort:          getEnv("DB_PORT", "3306"),
		DbUser:          getEnv("DB_USER", "root"),
		DbPass:          getEnv("DB_PASS", "root"),
		ServerName:      getEnv("SERVER_NAME", "127.0.0.1"),
		DbType:          getEnv("DB_TYPE", "sqlite"),
		ResponseVersion: int(getEnvAsInt("RESPONSE_VERSION", 1)),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, exists := os.LookupEnv(key); exists {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}

	return fallback
}
