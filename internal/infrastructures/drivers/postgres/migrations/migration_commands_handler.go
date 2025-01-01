package migrations

import (
	"fmt"
	"net/url"
	"strings"

	infrastructures_drivers_postgres "meul/inventory/internal/infrastructures/drivers/postgres"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

type MigrationPath string
type Command string
type Steps int
type ForceVersion int
type BaseName string
type MigrationConfig struct {
	MigrationPath MigrationPath
	Command       Command
	Steps         Steps
	ForceVersion  ForceVersion
	BaseName      BaseName
}

type MigrationHandler struct {
	DbConfig              infrastructures_drivers_postgres.DbConfig
	DB                    *gorm.DB
	MigrationConfig       *MigrationConfig
	MigrationFilesHandler *MigrationFilesHandler
}

func DefaultMigrationHandler(
	dbConfig infrastructures_drivers_postgres.DbConfig,
	db *gorm.DB,
	config *MigrationConfig,
	filesHandler *MigrationFilesHandler,
) *MigrationHandler {
	return &MigrationHandler{
		DbConfig:              dbConfig,
		DB:                    db,
		MigrationConfig:       config,
		MigrationFilesHandler: filesHandler,
	}
}

// initialize migration by converting the DSN to a connection string
func (m *MigrationHandler) initializeMigration() (*migrate.Migrate, error) {
	var connectionStringBuilder strings.Builder
	dsnParts := strings.Fields(m.DbConfig.DSN)
	dsnMap := make(map[string]string)

	for _, part := range dsnParts {
		kv := strings.SplitN(part, "=", 2)
		if len(kv) == 2 {
			dsnMap[kv[0]] = kv[1]
		}
	}

	if user, ok := dsnMap["user"]; ok {
		connectionStringBuilder.WriteString("postgres://" + user)
	}
	if password, ok := dsnMap["password"]; ok {
		connectionStringBuilder.WriteString(":" + password + "@")
	}
	if host, ok := dsnMap["host"]; ok {
		connectionStringBuilder.WriteString(host)
	}
	if port, ok := dsnMap["port"]; ok {
		connectionStringBuilder.WriteString(":" + port)
	}
	if dbname, ok := dsnMap["dbname"]; ok {
		connectionStringBuilder.WriteString("/" + dbname)
	}

	options := url.Values{}

	for key, value := range dsnMap {

		if key != "user" && key != "password" && key != "host" && key != "port" && key != "dbname" {
			options.Add(key, value)
		}
	}

	if encodedOptions := options.Encode(); encodedOptions != "" {
		connectionStringBuilder.WriteString("?" + encodedOptions)
	}
	connectionString := connectionStringBuilder.String()

	return migrate.New(string(m.MigrationConfig.MigrationPath), connectionString)
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

func (m *MigrationHandler) AutoMigrate(dst ...interface{}) error {
	err := m.DB.AutoMigrate(dst...)
	if err != nil {
		return fmt.Errorf("failed to AutoMigrate models: %w", err)
	}

	outputFiles, err := m.CatchMigrationsToSQLFiles()

	if err != nil {
		return fmt.Errorf("failed to catch migrations to sql files: %w", err)
	}

	for i := range outputFiles {

		if err := m.RunRenameFromString(outputFiles[i]); err != nil {
			return fmt.Errorf("failed to rename files from string: %w", err)
		}
	}
	fmt.Println("AutoMigration completed successfully.")
	return nil
}

func (m *MigrationHandler) RunStepsDown() error {
	return m.executeMigration(func(migration *migrate.Migrate) error {
		if err := migration.Steps(-int(m.MigrationConfig.Steps)); err != nil {
			return fmt.Errorf("failed to rollback migrations: %w", err)
		}
		fmt.Printf("Rolled back %d steps successfully\n", m.MigrationConfig.Steps)
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
	if err := m.MigrationFilesHandler.RenameMigrationFiles(string(m.MigrationConfig.BaseName)); err != nil {
		return fmt.Errorf("failed to rename migration files: %w", err)
	}
	return nil
}

func (m *MigrationHandler) RunRenameFromString(baseName string) error {
	if err := m.MigrationFilesHandler.RenameMigrationFiles(strings.TrimSuffix(baseName, ".sql")); err != nil {
		return fmt.Errorf("failed to rename migration files: %w", err)
	}
	return nil
}

func (m *MigrationHandler) RunForce() error {
	return m.executeMigration(func(migration *migrate.Migrate) error {
		if err := migration.Force(int(m.MigrationConfig.ForceVersion)); err != nil {
			return fmt.Errorf("failed to force migration version: %w", err)
		}
		fmt.Printf("Forced migration to version %d\n", m.MigrationConfig.ForceVersion)
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
