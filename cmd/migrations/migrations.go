package main

import (
	"log"
)

func main() {
	// Initialize the migration handler using Wire
	migrationHandler, err := InitializeMigrationHandler()
	if err != nil {
		log.Fatalf("failed to initialize migration handler: %v", err)
	}

	// Run the migration handler
	migrationHandler.Run()
}
