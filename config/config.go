package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db Database
}

// App - stores configuration
var App *Config

func Init() error {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load environment variables & populate App
	App = &Config{
		Db: Database{
			URL:  os.Getenv("MONGO_DB_URL"),
			Name: os.Getenv("MONGO_DB_NAME"),
		},
	}

	return nil

}
