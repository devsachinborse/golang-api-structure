package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	MongoURI   string
	Database   string
}

func LoadConfig() *Config {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Default value
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is required")
	}

	database := os.Getenv("MONGO_DB")
	if database == "" {
		database = "restaurant" // Default value
	}

	return &Config{
		ServerPort: port,
		MongoURI:   mongoURI,
		Database:   database,
	}
}
