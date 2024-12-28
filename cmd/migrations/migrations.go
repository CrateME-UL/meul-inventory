package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	migrations_resource "meul/inventory/internal/interfaces/cli"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// MigrationHandler encapsulates migration logic
type MigrationHandler struct {
	flags migrations_resource.MigrationFlags
}

func NewMigrationHandler(DatabaseURL string,
	MigrationsPath string,
	Command string,
	Steps int,
	ForceVersion int,
	BaseName string) MigrationHandler {

	flags := migrations_resource.MigrationFlags{
		DatabaseURL:    DatabaseURL,
		MigrationsPath: MigrationsPath,
		Command:        Command,
		Steps:          Steps,
		ForceVersion:   ForceVersion,
		BaseName:       BaseName,
	}

	migration_handler := MigrationHandler{
		flags,
	}

	return migration_handler
}

// Run executes the migration command
func (m *MigrationHandler) Run() {
	switch m.flags.Command {
	case "up":
		m.runUp()
	case "down":
		m.runDown()
	case "rename":
		m.runRename()
	case "force":
		m.runForce()
	case "version":
		m.runVersion()
	default:
		log.Fatalf("Invalid command: %s. Use 'up', 'down', 'force', 'version', or 'rename'", m.flags.Command)
	}
}

func (m *MigrationHandler) runUp() {
	migration := m.initializeMigration()
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	fmt.Println("Migrations applied successfully")
}

func (m *MigrationHandler) runDown() {
	if m.flags.Steps > 0 {
		migration := m.initializeMigration()
		if err := migration.Steps(-m.flags.Steps); err != nil {
			log.Fatalf("Failed to rollback migrations: %v", err)
		}
		fmt.Printf("Rolled back %d steps successfully\n", m.flags.Steps)
	} else {
		log.Fatal("For 'down', steps must be a positive number")
	}
}

func (m *MigrationHandler) runRename() {
	if m.flags.BaseName == "" {
		log.Fatal("You must provide a base name using the -base flag")
	}
	renameMigrationFiles(m.flags.BaseName)
}

func (m *MigrationHandler) runForce() {
	migration := m.initializeMigration()
	if err := migration.Force(m.flags.ForceVersion); err != nil {
		log.Fatalf("Failed to force migration version: %v", err)
	}
	fmt.Printf("Forced migration to version %d\n", m.flags.ForceVersion)
}

func (m *MigrationHandler) runVersion() {
	migration := m.initializeMigration()
	version, dirty, err := migration.Version()
	if err != nil {
		log.Fatalf("Failed to get migration version: %v", err)
	}
	fmt.Printf("Current version: %d, dirty: %v\n", version, dirty)
}

func (m *MigrationHandler) initializeMigration() *migrate.Migrate {
	if m.flags.DatabaseURL == "" && m.flags.Command != "rename" {
		log.Fatal("DATABASE_URL must be set or provided via the -db flag")
	}

	migration, err := migrate.New(m.flags.MigrationsPath, m.flags.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize migrate: %v", err)
	}
	return migration
}

func renameMigrationFiles(baseName string) {
	dir := filepath.Dir(baseName)
	name := filepath.Base(baseName)

	timestamp := time.Now().Format("20060102150405")
	orderPrefix := fmt.Sprintf("%04d", getNextMigrationOrder(dir))

	upFile := filepath.Join(dir, fmt.Sprintf("%s.up.sql", name))
	downFile := filepath.Join(dir, fmt.Sprintf("%s.down.sql", name))

	newUpFile := filepath.Join(dir, fmt.Sprintf("%s_%s_%s.up.sql", orderPrefix, timestamp, name))
	newDownFile := filepath.Join(dir, fmt.Sprintf("%s_%s_%s.down.sql", orderPrefix, timestamp, name))

	fmt.Printf("Looking for files:\n - %s\n - %s\n", upFile, downFile)
	fmt.Printf("Renaming to:\n - %s\n - %s\n", newUpFile, newDownFile)

	if err := os.Rename(upFile, newUpFile); err != nil {
		log.Fatalf("Failed to rename .up.sql file: %v", err)
	}
	if err := os.Rename(downFile, newDownFile); err != nil {
		log.Fatalf("Failed to rename .down.sql file: %v", err)
	}

	fmt.Printf("Files renamed successfully:\n")
	fmt.Printf("Old .up.sql: %s -> New .up.sql: %s\n", upFile, newUpFile)
	fmt.Printf("Old .down.sql: %s -> New .down.sql: %s\n", downFile, newDownFile)
}

func getNextMigrationOrder(dir string) int {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("Failed to read migration directory: %v", err)
	}

	var highestOrder int = -1
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		matches := regexp.MustCompile(`^(\d{4})_\d{14}_.+\.up\.sql$`).FindStringSubmatch(file.Name())
		if len(matches) > 0 {
			order, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Printf("Invalid migration order in file %s: %v", file.Name(), err)
				continue
			}

			if order > highestOrder {
				highestOrder = order
			}
		}
	}

	return highestOrder + 1
}

func main() {
	flags := migrations_resource.NewMigrationFlags()
	migration_handler := NewMigrationHandler(
		flags.DatabaseURL, flags.MigrationsPath, flags.Command,
		flags.Steps, flags.ForceVersion, flags.BaseName)
	migration_handler.Run()
}
