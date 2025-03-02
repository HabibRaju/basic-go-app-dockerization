package main

import (
	"log"
	"net/http"
	"sample-health/internal/config"
	"sample-health/internal/routes"
	"sample-health/pkg/database"
)

func main() {

	cfg := config.LoadConfig()

	// Initialize database connection
	db, err := database.NewMySQLConnection(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Set up routes
	router := routes.SetupRoutes(db)

	// Start the server
	log.Println("Server is running on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
