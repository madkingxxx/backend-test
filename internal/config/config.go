package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort int

	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string

	SkinportAPIBaseURL string
	DBSSLMode          string

	ItemsCronExpression string
	LogLevel            string
}

// New - creates a new Config with default values.
func New() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("no .env file found")
	}

	viper.AutomaticEnv()
	cfg := &Config{
		ServerPort:          viper.GetInt("SERVER_PORT"),
		DBHost:              viper.GetString("DB_HOST"),
		DBPort:              viper.GetInt("DB_PORT"),
		DBUser:              viper.GetString("DB_USER"),
		DBPassword:          viper.GetString("DB_PASSWORD"),
		DBName:              viper.GetString("DB_NAME"),
		SkinportAPIBaseURL:  viper.GetString("SKINPORT_API_BASE_URL"),
		DBSSLMode:           viper.GetString("DB_SSL_MODE"),
		ItemsCronExpression: viper.GetString("ITEMS_CRON_EXPRESSION"),
		LogLevel:            viper.GetString("LOG_LEVEL"),
	}

	return cfg
}
