package postgres_migrations

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
)

// MigrationHandler holds all the commands and their values
type MigrationHandler struct {
	DatabaseURL    string
	MigrationsPath string
	Command        string
	Steps          int
	ForceVersion   int
	BaseName       string
}

func (m *MigrationHandler) RunUp() {
	migration := m.initializeMigration()
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	fmt.Println("Migrations applied successfully")
}

func (m *MigrationHandler) RunDown() {
	if m.Steps > 0 {
		migration := m.initializeMigration()
		if err := migration.Steps(-m.Steps); err != nil {
			log.Fatalf("Failed to rollback migrations: %v", err)
		}
		fmt.Printf("Rolled back %d steps successfully\n", m.Steps)
	} else {
		log.Fatal("For 'down', steps must be a positive number")
	}
}

func (m *MigrationHandler) RunRename() {
	if m.BaseName == "" {
		log.Fatal("You must provide a base name using the -base flag")
	}
	renameMigrationFiles(m.BaseName)
}

func (m *MigrationHandler) RunForce() {
	migration := m.initializeMigration()
	if err := migration.Force(m.ForceVersion); err != nil {
		log.Fatalf("Failed to force migration version: %v", err)
	}
	fmt.Printf("Forced migration to version %d\n", m.ForceVersion)
}

func (m *MigrationHandler) RunVersion() {
	migration := m.initializeMigration()
	version, dirty, err := migration.Version()
	if err != nil {
		log.Fatalf("Failed to get migration version: %v", err)
	}
	fmt.Printf("Current version: %d, dirty: %v\n", version, dirty)
}

func (m *MigrationHandler) initializeMigration() *migrate.Migrate {
	if m.DatabaseURL == "" && m.Command != "rename" {
		log.Fatal("DATABASE_URL must be set or provided via the -db flag")
	}

	migration, err := migrate.New(m.MigrationsPath, m.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize migrate: %v", err)
	}
	return migration
}
