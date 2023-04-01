package link

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Link() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	tgToken := os.Getenv("t")
	URL := "https://api.telegram.org/bot" + tgToken
	return URL
}
