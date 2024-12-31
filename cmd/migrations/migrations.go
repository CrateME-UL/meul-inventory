package main

import (
	"log"
	"meul/inventory/internal/infrastructures/drivers/postgres/models"
)

func main() {
	// Initialize the migration handler
	migrationHandler, err := InitializeMigrationHandler()
	if err != nil {
		log.Fatalf("failed to initialize migration handler: %v", err)
	}

	// Auto-migrate the models
	err = migrationHandler.DB.AutoMigrate(&models.Place{}, &models.Item{}, &models.PlaceItems{})
	if err != nil {
		log.Fatalf("failed to auto-migrate models: %v", err)
	}

	// Run the migration handler
	// migrationHandler.Run()

	log.Println("Database migration completed successfully.")
}
