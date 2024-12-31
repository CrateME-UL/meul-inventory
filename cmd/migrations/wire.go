//go:build wireinject
// +build wireinject

package main

import (
	postgres_migrations "meul/inventory/internal/infrastructures/drivers/postgres/migrations"
	migrations_resource "meul/inventory/internal/interfaces/migrations_resource_cli"

	"github.com/google/wire"
)

func ProvideMigrationHandler(config *postgres_migrations.MigrationConfig) postgres_migrations.MigrationHandler {
	// Create the handler based on the config
	return postgres_migrations.MigrationHandler{
		MigrationConfig: config,
	}
}

// WireSet is a set that includes all necessary providers
var WireSet = wire.NewSet(
	// Existing providers
	migrations_resource.DefaultMigrationCLI,
	postgres_migrations.DefaultMigrationFilesOrderHandler,
	postgres_migrations.DefaultMigrationFilesHandler,
	postgres_migrations.DefaultMigrationHandler,
	// Wire the MigrationConfig from the global variable
	wire.Value(&migrations_resource.MigrationConfig),
	ProvideMigrationHandler,
)

func InitializeMigrationHandler() (*migrations_resource.MigrationCLI, error) {
	wire.Build(WireSet)

	return nil, nil
}
