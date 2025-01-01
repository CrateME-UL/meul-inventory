package main

import (
	"log"
)

func main() {
	cli, err := InitializeMigrationHandler()
	if err != nil {
		log.Fatalf("failed to initialize migration handler: %v", err)
	}

	cli.Run()

	log.Println("Database migration completed successfully.")
}
