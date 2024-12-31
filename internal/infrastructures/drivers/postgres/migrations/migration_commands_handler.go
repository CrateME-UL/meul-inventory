package postgres_migrations

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// MigrationHandler holds all the commands and their values
type MigrationHandler struct {
	DatabaseURL           string
	MigrationsPath        string
	Command               string
	Steps                 int
	ForceVersion          int
	BaseName              string
	MigrationFilesHandler *MigrationFilesHandler
}

func (m *MigrationHandler) initializeMigration() (*migrate.Migrate, error) {
	return migrate.New(m.MigrationsPath, m.DatabaseURL)
}

func (m *MigrationHandler) executeMigration(action func(*migrate.Migrate) error) error {
	migration, err := m.initializeMigration()
	if err != nil {
		return fmt.Errorf("failed to initialize migrate: %w", err)
	}
	defer migration.Close()

	if err := action(migration); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration error: %w", err)
	}

	return nil
}

func (m *MigrationHandler) RunUp() error {
	return m.executeMigration(func(migration *migrate.Migrate) error {
		if err := migration.Up(); err != nil {
			return fmt.Errorf("failed to apply migrations: %w", err)
		}
		fmt.Println("Migrations applied successfully")
		return nil
	})
}

func (m *MigrationHandler) RunStepsDown() error {
	return m.executeMigration(func(migration *migrate.Migrate) error {
		if err := migration.Steps(-m.Steps); err != nil {
			return fmt.Errorf("failed to rollback migrations: %w", err)
		}
		fmt.Printf("Rolled back %d steps successfully\n", m.Steps)
		return nil
	})
}

func (m *MigrationHandler) RunDown() error {
	return m.executeMigration(func(migration *migrate.Migrate) error {
		if err := migration.Down(); err != nil {
			return fmt.Errorf("failed to rollback all migrations: %w", err)
		}
		fmt.Printf("Rolled back all migrations successfully\n")
		return nil
	})
}

func (m *MigrationHandler) RunRename() error {
	if err := m.MigrationFilesHandler.RenameMigrationFiles(m.BaseName); err != nil {
		return fmt.Errorf("failed to rename migration files: %w", err)
	}
	return nil
}

func (m *MigrationHandler) RunForce() error {
	return m.executeMigration(func(migration *migrate.Migrate) error {
		if err := migration.Force(m.ForceVersion); err != nil {
			return fmt.Errorf("failed to force migration version: %w", err)
		}
		fmt.Printf("Forced migration to version %d\n", m.ForceVersion)
		return nil
	})
}

func (m *MigrationHandler) RunVersion() error {
	return m.executeMigration(func(migration *migrate.Migrate) error {
		version, dirty, err := migration.Version()
		if err != nil {
			return fmt.Errorf("failed to get migration version: %w", err)
		}
		fmt.Printf("Current version: %d, dirty: %v\n", version, dirty)
		return nil
	})
}
