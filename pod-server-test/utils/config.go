package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MapsKey         string
	FirebaseKeyPath string
}

func LoadConfig() Config {
	// Load .env file if it exists, otherwise rely on system environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	return Config{
		MapsKey:         os.Getenv("GOOGLE_MAPS_API_KEY"),
		FirebaseKeyPath: os.Getenv("FIREBASE_CREDENTIALS"),
	}
}
