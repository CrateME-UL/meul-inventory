package migrations_resource

import (
	"flag"
	"log"
	postgres_migrations "meul/inventory/internal/infrastructures/drivers/postgres/migrations"
	"os"
)

// NewMigrationFlags parses and returns MigrationFlags
func NewMigrationFlags() postgres_migrations.MigrationHandler {
	flags := postgres_migrations.MigrationHandler{}
	flag.StringVar(&flags.DatabaseURL, "db", os.Getenv("DATABASE_URL"), "Database connection URL")
	flag.StringVar(&flags.MigrationsPath, "path", "file://migrations", "Path to migrations directory")
	flag.StringVar(&flags.Command, "command", "up", "Migration command: up, down, force, version, or rename")
	flag.IntVar(&flags.Steps, "steps", 1, "Number of steps for the 'down' command")
	flag.IntVar(&flags.ForceVersion, "version", 0, "Version to force with the 'force' command")
	flag.StringVar(&flags.BaseName, "base", "", "Base name of the migration file (without '.up.sql' or '.down.sql')")
	flag.Parse()
	return flags
}

// MigrationHandler encapsulates migration logic
type MigrationHandler struct {
	flags postgres_migrations.MigrationHandler
}

// Run executes the migration command
func (m *MigrationHandler) Run() {
	switch m.flags.Command {
	case "up":
		m.flags.RunUp()
	case "down":
		m.flags.RunDown()
	case "rename":
		m.flags.RunRename()
	case "force":
		m.flags.RunForce()
	case "version":
		m.flags.RunVersion()
	default:
		log.Fatalf("Invalid command: %s. Use 'up', 'down', 'force', 'version', or 'rename'", m.flags.Command)
	}
}

func NewMigrationHandler(DatabaseURL string,
	MigrationsPath string,
	Command string,
	Steps int,
	ForceVersion int,
	BaseName string) MigrationHandler {

	flags := postgres_migrations.MigrationHandler{
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
