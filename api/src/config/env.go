package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBName          = "devbook"
	DBUser          = "devbook"
	DBPassword      = "devbook"
	DBHost          = "127.0.0.1"
	DBPort          = "3300"
	APIPort         = "4000"
	APIJWTSignature = "secret"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBName = os.Getenv("DB_NAME")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	APIPort = os.Getenv("API_PORT")
	APIJWTSignature = os.Getenv("API_JWT_TOKEN=secret")
}
