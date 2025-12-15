package config

import (
	"ShoppingCartAnalysis/internal/logger"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	PG_Port string
	PG_Host string
	PG_User string
	PG_Password string
	PG_Dbname string
	GeminiAPIKey string
}

func Load() *Config {
	logger.LoggerInfo.Println("Initializing config")

	err := godotenv.Load()

	if err != nil {
		logger.LoggerInfo.Fatalln("Error loading env file: ", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	pgPort := os.Getenv("PG_PORT")
	pgHost := os.Getenv("PG_HOST")
	pgUser := os.Getenv("PG_USERNAME")
	pgPassword := os.Getenv("PG_DBPASSWORD")
	pgDbname := os.Getenv("PG_DBNAME")

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		logger.LoggerError.Fatalln("GEMINI_API_KEY not set in environment or .env file")
	}

	logger.LoggerInfo.Println("Initialization config completed")

	return &Config{
		Port: port,
		PG_Port: pgPort,
		PG_Host: pgHost,
		PG_User: pgUser,
		PG_Password: pgPassword,
		PG_Dbname: pgDbname,
		GeminiAPIKey: apiKey,
	}
}