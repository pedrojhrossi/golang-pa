package config

import (
	"fmt"
	"os"
)

type Config struct {
	MongoURI string
	DBName   string
	Port     string
}

func LoadConfig() *Config {
	dbUser := getEnv("MONGO_USER", "admin")
	dbPass := getEnv("MONGO_PASS", "12345")
	dbHost := getEnv("MONGO_HOST", "localhost")
	dbPort := getEnv("MONGO_PORT", "27017")
	dbName := getEnv("DB_NAME", "appointment_db")
	appPort := getEnv("APP_PORT", "9080")

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	return &Config{
		MongoURI: mongoURI,
		DBName:   dbName,
		Port:     appPort,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
