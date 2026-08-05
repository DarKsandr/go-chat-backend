package pkg

import (
	"log"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}
