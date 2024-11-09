package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBURL string

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	os.Getenv("DBURL")

}
