package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	GoogleAPIKey string
	GoogleCX     string
	GoogleURL    string
	DailyLimit   = 100             // Google API free tier limit
	CounterFile  = "consultas.json" // File to track queries
	APITimeout   = 15 * time.Second // API request timeout
)

// LoadEnv loads environment variables from .env file
func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	GoogleAPIKey = os.Getenv("GOOGLE_API_KEY")
	GoogleCX = os.Getenv("GOOGLE_CX")
	GoogleURL = os.Getenv("GOOGLE_URL")
	if GoogleAPIKey == "" || GoogleCX == "" || GoogleURL == "" {
		return fmt.Errorf("required environment variables GOOGLE_API_KEY, GOOGLE_CX, and GOOGLE_URL must be set")
	}
	return nil
}