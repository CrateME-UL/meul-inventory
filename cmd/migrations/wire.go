//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
	infrastructures_drivers_postgres "meul/inventory/internal/infrastructures/drivers/postgres"
	postgres_migrations "meul/inventory/internal/infrastructures/drivers/postgres/migrations"
	migrations_resource "meul/inventory/internal/interfaces/migrations_resource_cli"

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

// ProvideMigrationHandler creates a MigrationHandler based on the provided DB and config
// func ProvideMigrationHandler(db *gorm.DB, config *postgres_migrations.MigrationConfig) *postgres_migrations.MigrationHandler {
// 	return &postgres_migrations.MigrationHandler{
// 		DB:              db,
// 		MigrationConfig: config,
// 	}
// }

// WireSet is a set that includes all necessary providers
var WireSet = wire.NewSet(
	// Existing providers
	ProvideConfig,
	infrastructures_drivers_postgres.NewDatabaseConnection,
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
