package config

import "os"

type AppConfiguration struct {
	AppPort      string
	AppEnv       string
	JWTSecret    string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPass       string
	DBName       string
	DBSSLMode    string
	DBTimeZone   string
	RedisHost    string
	RedisPort    string
	RedisPass    string
	RabbitMQHost string
	RabbitMQPort string
	RabbitMQUser string
	RabbitMQPass string
}

var AppConfig AppConfiguration

func LoadConfig() {
	AppConfig = AppConfiguration{
		AppPort:      os.Getenv("PORT"),
		AppEnv:       os.Getenv("APP_ENV"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPass:       os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		DBSSLMode:    os.Getenv("DB_SSLMODE"),
		DBTimeZone:   os.Getenv("DB_TIMEZONE"),
		RedisHost:    os.Getenv("REDIS_HOST"),
		RedisPort:    os.Getenv("REDIS_PORT"),
		RedisPass:    os.Getenv("REDIS_PASS"),
		RabbitMQHost: os.Getenv("RABBITMQ_HOST"),
		RabbitMQPort: os.Getenv("RABBITMQ_PORT"),
		RabbitMQUser: os.Getenv("RABBITMQ_USER"),
		RabbitMQPass: os.Getenv("RABBITMQ_PASS"),
	}
}
