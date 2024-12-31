// cli for migrations
package migrations_resource_cli

import (
	"flag"
	"fmt"
	"log"
	postgres_migrations "meul/inventory/internal/infrastructures/drivers/postgres/migrations"
	"os"
	"strings"
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

	flags.BaseName = strings.Trim(flags.BaseName, "")
	flags.Command = strings.Trim(flags.Command, "")
	flags.MigrationsPath = strings.Trim(flags.MigrationsPath, "")
	flags.DatabaseURL = strings.Trim(flags.DatabaseURL, "")

	return flags
}

// MigrationHandler encapsulates migration logic
type MigrationHandler struct {
	flags postgres_migrations.MigrationHandler
}

// Run executes the migration command
func (m *MigrationHandler) Run() {
	var err error

	switch m.flags.Command {
	case "up":
		err = validateDbFlag(m)

		if err == nil {
			err = m.flags.RunUp()
		}
	case "down":
		err = validateDbFlag(m)
		if err == nil && m.flags.Steps <= 0 {
			err = fmt.Errorf("for 'down', steps must be a positive number: %v", err)
		}
		if err == nil {
			err = m.flags.RunStepsDown()
		}
	case "rename":
		if m.flags.BaseName == "" {
			err = fmt.Errorf("you must provide a base name using the -base flag: %v", err)
		}
		if err == nil {
			err = m.flags.RunRename()
		}
	case "force":
		err = validateDbFlag(m)

		if err == nil {
			err = m.flags.RunForce()
		}
	case "version":
		err = validateDbFlag(m)

		if err == nil {
			err = m.flags.RunVersion()
		}
	default:
		log.Fatalf("Invalid command: %s. Use 'up', 'down', 'force', 'version', or 'rename'", m.flags.Command)
	}

	if err != nil {
		log.Fatal(err)
	}

}

func validateDbFlag(m *MigrationHandler) (err error) {

	if m.flags.DatabaseURL == "" {

		return fmt.Errorf("DATABASE_URL must be set via the -db flag %v", err)
	}

	return err
}

// Create a MigrationHandler from flags
func NewMigrationHandler(DatabaseURL string,
	MigrationsPath string,
	Command string,
	Steps int,
	ForceVersion int,
	BaseName string,
	MigrationFilesHandler *postgres_migrations.MigrationFilesHandler) MigrationHandler {

	flags := postgres_migrations.MigrationHandler{
		DatabaseURL:           DatabaseURL,
		MigrationsPath:        MigrationsPath,
		Command:               Command,
		Steps:                 Steps,
		ForceVersion:          ForceVersion,
		BaseName:              BaseName,
		MigrationFilesHandler: MigrationFilesHandler,
	}

	migration_handler := MigrationHandler{
		flags,
	}

	return migration_handler
}
