package main

import (
	postgres_migrations "meul/inventory/internal/infrastructures/drivers/postgres/migrations"
	migrations_resource "meul/inventory/internal/interfaces/migrations_resource_cli"
)

// todo: add wire setup
func main() {
	flags := migrations_resource.NewMigrationFlags()

	migrationFileOrderHandler := postgres_migrations.NewMigrationFilesOrderHandler()
	migrationFileHandler := postgres_migrations.NewMigrationFilesHandler(migrationFileOrderHandler)

	migration_handler := migrations_resource.NewMigrationHandler(
		flags.DatabaseURL, flags.MigrationsPath, flags.Command,
		flags.Steps, flags.ForceVersion, flags.BaseName, migrationFileHandler)
	migration_handler.Run()
}
