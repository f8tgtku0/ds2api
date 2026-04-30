package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Version is set at build time via ldflags
var Version = "dev"

func main() {
	// Load .env file if present (ignored in production where env vars are set directly)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Printf("Starting ds2api %s on port %s", Version, cfg.Port)
	log.Printf("Build info: version=%s", Version)
	// Log the Go environment for easier debugging during local development
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development" // default to development if APP_ENV is not set
	}
	log.Printf("Environment: %s", appEnv)

	server := NewServer(cfg)
	if err := server.Run(); err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}
}
