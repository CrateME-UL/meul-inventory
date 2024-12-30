package postgres_migrations

import (
	"fmt"

	"github.com/golang-migrate/migrate"
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

func (m *MigrationHandler) RunUp() (err error) {
	var migration *migrate.Migrate

	if migration, err = m.initializeMigration(); err != nil {

		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {

		return fmt.Errorf("failed to apply migrations: %w", err)
	}
	fmt.Println("Migrations applied successfully")

	return err
}

func (m *MigrationHandler) RunDown() (err error) {
	var migration *migrate.Migrate

	if migration, err = m.initializeMigration(); err != nil {

		return fmt.Errorf("failed to rollback migrations: %w", err)
	}

	if err := migration.Steps(-m.Steps); err != nil {

		return fmt.Errorf("failed to rollback migrations: %w", err)
	}
	fmt.Printf("Rolled back %d steps successfully\n", m.Steps)

	return err
}

func (m *MigrationHandler) RunRename() (err error) {
	if err := m.MigrationFilesHandler.renameMigrationFiles(m.BaseName); err != nil {

		return fmt.Errorf("failed to rename migration files: %w", err)
	}

	return err
}

func (m *MigrationHandler) RunForce() (err error) {
	var migration *migrate.Migrate

	if migration, err = m.initializeMigration(); err != nil {

		return fmt.Errorf("failed to force migration version: %w", err)
	}

	if err := migration.Force(m.ForceVersion); err != nil {

		return fmt.Errorf("failed to force migration version: %w", err)
	}
	fmt.Printf("Forced migration to version %d\n", m.ForceVersion)

	return err
}

func (m *MigrationHandler) RunVersion() (err error) {
	var version uint
	var dirty bool
	var migration *migrate.Migrate

	if migration, err = m.initializeMigration(); err != nil {

		return fmt.Errorf("failed to get migration version: %w", err)
	}

	if version, dirty, err = migration.Version(); err != nil {

		return fmt.Errorf("failed to get migration version: %w", err)
	}
	fmt.Printf("Current version: %d, dirty: %v\n", version, dirty)

	return err
}

func (m *MigrationHandler) initializeMigration() (migration *migrate.Migrate, err error) {
	migration, err = migrate.New(m.MigrationsPath, m.DatabaseURL)

	if err != nil {

		return nil, fmt.Errorf("failed to initialize migrate: %w: ", err)
	}

	return migration, err
}
