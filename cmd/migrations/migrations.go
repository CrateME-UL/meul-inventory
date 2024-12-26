package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Command-line flags
	databaseURL := flag.String("db", os.Getenv("DATABASE_URL"), "Database connection URL")
	migrationsPath := flag.String("path", "file://migrations", "Path to migrations directory")
	command := flag.String("command", "up", "Migration command: up, down, force, version, or create")
	steps := flag.Int("steps", 1, "Number of steps for the 'down' command")
	forceVersion := flag.Int("version", 0, "Version to force with the 'force' command")
	baseName := flag.String("base", "", "Base name of the migration file (without '.up.sql' or '.down.sql')")
	flag.Parse()

	switch *command {
	case "up":
		m := initializeMigration(databaseURL, command, migrationsPath)
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
		fmt.Println("Migrations applied successfully")
	case "down":
		if *steps > 0 {
			m := initializeMigration(databaseURL, command, migrationsPath)
			if err := m.Steps(-*steps); err != nil {
				log.Fatalf("Failed to rollback migrations: %v", err)
			}
			fmt.Printf("Rolled back %d steps successfully\n", *steps)
		} else {
			log.Fatal("For 'down', steps must be a positive number")
		}
	case "rename":
		if *baseName == "" {
			log.Fatal("You must provide a base name using the -base flag")
		}
		renameMigrationFiles(*baseName)
	case "force":
		m := initializeMigration(databaseURL, command, migrationsPath)
		if err := m.Force(*forceVersion); err != nil {
			log.Fatalf("Failed to force migration version: %v", err)
		}
		fmt.Printf("Forced migration to version %d\n", *forceVersion)
	case "version":
		m := initializeMigration(databaseURL, command, migrationsPath)
		version, dirty, err := m.Version()
		if err != nil {
			log.Fatalf("Failed to get migration version: %v", err)
		}
		fmt.Printf("Current version: %d, dirty: %v\n", version, dirty)
	default:
		log.Fatalf("Invalid command: %s. Use 'up', 'down', 'force', 'version', or 'create'", *command)
	}
}

func initializeMigration(databaseURL *string, command *string, migrationsPath *string) *migrate.Migrate {
	if *databaseURL == "" && *command != "rename" {
		log.Fatal("DATABASE_URL must be set or provided via the -db flag")
	}

	// Initialize migration
	m, err := migrate.New(*migrationsPath, *databaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize migrate: %v", err)
	}
	return m
}

func renameMigrationFiles(baseName string) {
	dir := filepath.Dir(baseName)
	name := filepath.Base(baseName)

	// Get the current timestamp in the format YYYYMMDDHHMMSS
	timestamp := time.Now().Format("20060102150405")

	// Find the next available migration order
	orderPrefix := fmt.Sprintf("%04d", getNextMigrationOrder(dir)) // Pad with leading zeros to make it 4 digits (0001, 0002, etc.)

	// File paths for .up.sql and .down.sql
	upFile := filepath.Join(dir, fmt.Sprintf("%s.up.sql", name))
	downFile := filepath.Join(dir, fmt.Sprintf("%s.down.sql", name))

	// New file names with the order prefix
	newUpFile := filepath.Join(dir, fmt.Sprintf("%s_%s_%s.up.sql", orderPrefix, timestamp, name))
	newDownFile := filepath.Join(dir, fmt.Sprintf("%s_%s_%s.down.sql", orderPrefix, timestamp, name))

	// Print paths for debugging
	fmt.Printf("Looking for files:\n - %s\n - %s\n", upFile, downFile)

	// Debug: Print new file paths
	fmt.Printf("Renaming to:\n - %s\n - %s\n", newUpFile, newDownFile)

	// Rename files
	if err := os.Rename(upFile, newUpFile); err != nil {
		log.Fatalf("Failed to rename .up.sql file: %v", err)
	}
	if err := os.Rename(downFile, newDownFile); err != nil {
		log.Fatalf("Failed to rename .down.sql file: %v", err)
	}

	// Confirm renaming
	fmt.Printf("Files renamed successfully:\n")
	fmt.Printf("Old .up.sql: %s -> New .up.sql: %s\n", upFile, newUpFile)
	fmt.Printf("Old .down.sql: %s -> New .down.sql: %s\n", downFile, newDownFile)
}

// Function to get the next available migration order by finding the highest number
func getNextMigrationOrder(dir string) int {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("Failed to read migration directory: %v", err)
	}

	var highestOrder int
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Look for files with the pattern <order>_<timestamp>_name.up.sql
		matches := regexp.MustCompile(`^(\d{4})_\d{14}_.+\.up\.sql$`).FindStringSubmatch(file.Name())
		if len(matches) > 0 {
			order, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Printf("Invalid migration order in file %s: %v", file.Name(), err)
				continue
			}

			// Track the highest order number found
			if order > highestOrder {
				highestOrder = order
			}
		}
	}

	// Return the next order number, incremented by 1
	return highestOrder + 1
}
