//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
	infrastructures_drivers_postgres "meul/inventory/internal/infrastructures/drivers/postgres"
	postgres_migrations "meul/inventory/internal/infrastructures/drivers/postgres/migrations"
	migrations_resource "meul/inventory/internal/interfaces/migrations_resource_cli"
	"os"

	"github.com/google/wire"
)

var (
	buildMode    string
	port         string = ":3000"
	trustedProxy        = []string{"127.0.0.1", "::1"}
	dbHost       string
	dbUser       string
	dbPassword   string
	dbName       string
	dbPort       string
	dbSSLMode    string
	dbTimeZone   string = "America/Toronto"
)

// ProvideConfig creates a new gorm.DB instance for PostgreSQL
func ProvideConfig() (infrastructures_drivers_postgres.DbConfig, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode, dbTimeZone,
	)
	dbConfig := infrastructures_drivers_postgres.DbConfig{
		DSN: dsn,
	}

	return dbConfig, nil
}

func ProvideLogFile() (*os.File, error) {
	logFile, err := os.OpenFile("../internal/infrastructures/drivers/postgres/migrations/migration.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

// WireSet is a set that includes all necessary providers
var WireSet = wire.NewSet(
	// Existing providers
	ProvideConfig,
	ProvideLogFile,
	infrastructures_drivers_postgres.NewDatabaseConnectionWithMigrationLogger,
	migrations_resource.DefaultMigrationCLI,
	postgres_migrations.DefaultMigrationFilesOrderHandler,
	postgres_migrations.DefaultMigrationFilesHandler,
	postgres_migrations.DefaultMigrationHandler,
	// Wire the MigrationConfig from the global variable
	wire.Value(&migrations_resource.MigrationConfig),
	// ProvideMigrationHandler,
)

func InitializeMigrationHandler() (*postgres_migrations.MigrationHandler, error) {
	wire.Build(WireSet)

	return nil, nil
}
