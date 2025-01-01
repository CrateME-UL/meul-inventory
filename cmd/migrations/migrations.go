package main

import (
	"log"
	"meul/inventory/internal/infrastructures/drivers/postgres/migrations"
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

	outputFiles, err := migrations.CatchMigrationsToSQLFiles()

	if err != nil {
		log.Fatalf("failed to catch migrations to sql files: %v", err)
	}

	for i := range outputFiles {

		if err := migrationHandler.RunRenameFromString(outputFiles[i]); err != nil {
			log.Fatalf("failed to rename files from string: %v", err)
		}
	}

	// Run the migration handler
	migrationHandler.RunUp()

	log.Println("Database migration completed successfully.")
}
