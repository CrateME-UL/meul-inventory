package main

import (
	migrations_resource "meul/inventory/internal/interfaces/cli"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// todo: add wire setup
func main() {
	flags := migrations_resource.NewMigrationFlags()
	migration_handler := migrations_resource.NewMigrationHandler(
		flags.DatabaseURL, flags.MigrationsPath, flags.Command,
		flags.Steps, flags.ForceVersion, flags.BaseName)
	migration_handler.Run()
}
