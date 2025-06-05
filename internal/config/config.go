package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env файл не найден — пытаемся использовать переменные окружения")
	}

	return &Config{
		DBHost:     mustGetEnv("DB_HOST"),
		DBPort:     mustGetEnv("DB_PORT"),
		DBUser:     mustGetEnv("DB_USER"),
		DBPassword: mustGetEnv("DB_PASSWORD"),
		DBName:     mustGetEnv("DB_NAME"),
	}
}

func mustGetEnv(key string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("ОШИБКА: переменная окружения %s не установлена", key)
	}
	return val
}