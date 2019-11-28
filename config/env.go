package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() (string, string, string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	dbHost, exists := os.LookupEnv("DB_HOST")

	if !exists {
		log.Fatal("DB_HOST not defined in .env")
	}

	dbUser, exists := os.LookupEnv("DB_USER")

	if !exists {
		log.Fatal("DB_USER not defined in .env")
	}

	dbPass, exists := os.LookupEnv("DB_PASS")

	if !exists {
		log.Fatal("DB_PASS not defined in .env")
	}

	return dbHost, dbPass, dbUser
}
