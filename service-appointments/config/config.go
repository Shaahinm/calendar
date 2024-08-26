package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port      string
	DbName    string
	DbAddress string
	DbUser    string
	DbPass    string
}

var Envs = initConfig()

func initConfig() Config {
	// we can use the following line to read env vars from a file, but they won't be typed
	// we have to access them using string keys
	// godotenv.Load()
	//

	return Config{
		Port:      getEnv("PORT", "8080"),
		DbName:    getEnv("DB_NAME", "test.db"),
		DbAddress: getEnv("DB_ADDRESS", "localhost"),
		DbUser:    getEnv("DB_USER", "root"),
		DbPass:    getEnv("DB_PASS", "root"),
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
