package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBURL string
var JSONPATH string

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBURL = os.Getenv("DBURL")
	JSONPATH = os.Getenv("JSONPATH")
}
