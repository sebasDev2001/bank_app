package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(variable string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
	return os.Getenv(variable)
}
