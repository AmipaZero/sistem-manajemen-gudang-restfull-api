package util

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

var SecretKey string

func InitEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading system environment")
	}

	// Ambil SECRET_KEY dari env
	SecretKey = os.Getenv("SECRET_KEY")
	if SecretKey == "" {
		log.Fatal("SECRET_KEY is required but not set")
	}
}
