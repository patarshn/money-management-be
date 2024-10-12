package config

import (
	"os"
	"strconv"
)

type App struct {
	Name    string
	Env     string
	Version string
	Host    string
	Port    string
}

type Config struct {
	App App
}

func InitConfig() (*Config, error) {
	var config = Config{
		App: App{
			Name:    getEnv("APP_NAME", "money-management-be"),
			Env:     getEnv("APP_ENV", "development"),
			Version: getEnv("APP_VERSION", "1.0.0"),
			Host:    getEnv("APP_HOST", "localhost"),
			Port:    getEnv("APP_PORT", "8888"),
		},
	}

	return &config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvBool returns the boolean value of an environment variable or a default value
func getEnvBool(key string, defaultValue bool) bool {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}
	result, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return result
}

// getEnvInt returns the integer value of an environment variable or a default value
func getEnvInt(key string, defaultValue int) int {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return result
}

// getEnvFloat returns the float value of an environment variable or a default value
func getEnvFloat(key string, defaultValue float64) float64 {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue
	}
	return result
}
