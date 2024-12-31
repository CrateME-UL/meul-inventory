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

var MigrationConfig postgres_migrations.MigrationConfig

func init() {
	var dbURL string
	var path string
	var command string
	var steps int
	var forceVersion int
	var baseName string

	// Define command-line flags
	flag.StringVar(&dbURL, "db", os.Getenv("DATABASE_URL"), "Database connection URL")
	flag.StringVar(&path, "path", "file://migrations", "Path to migrations directory")
	flag.StringVar(&command, "command", "up", "Migration command: up, down, force, version, or rename")
	flag.IntVar(&steps, "steps", 1, "Number of steps for the 'down' command")
	flag.IntVar(&forceVersion, "version", 0, "Version to force with the 'force' command")
	flag.StringVar(&baseName, "base", "", "Base name of the migration file (without '.up.sql' or '.down.sql')")

	// Parse the flags
	flag.Parse()

	// Assign values to the global migrationConfig variable
	MigrationConfig = postgres_migrations.MigrationConfig{
		DatabaseURL:   postgres_migrations.DatabaseURL(strings.TrimSpace(dbURL)),
		MigrationPath: postgres_migrations.MigrationPath(strings.TrimSpace(path)),
		Command:       postgres_migrations.Command(strings.TrimSpace(command)),
		Steps:         postgres_migrations.Steps(steps),
		ForceVersion:  postgres_migrations.ForceVersion(forceVersion),
		BaseName:      postgres_migrations.BaseName(strings.TrimSpace(baseName)),
	}
}

// MigrationHandler encapsulates migration logic
type MigrationCLI struct {
	handler postgres_migrations.MigrationHandler
}

func DefaultMigrationCLI(handler *postgres_migrations.MigrationHandler) (cli *MigrationCLI) {
	cli = &MigrationCLI{
		handler: *handler,
	}
	return cli
}

// Run executes the migration command
func (m *MigrationCLI) Run() {
	var err error

	switch m.handler.MigrationConfig.Command {
	case "up":
		err = validateDbFlag(m)

		if err == nil {
			err = m.handler.RunUp()
		}
	case "down":
		err = validateDbFlag(m)
		if err == nil && m.handler.MigrationConfig.Steps <= 0 {
			err = fmt.Errorf("for 'down', steps must be a positive number: %v", err)
		}
		if err == nil {
			err = m.handler.RunStepsDown()
		}
	case "rename":
		if m.handler.MigrationConfig.BaseName == "" {
			err = fmt.Errorf("you must provide a base name using the -base flag: %v", err)
		}
		if err == nil {
			err = m.handler.RunRename()
		}
	case "force":
		err = validateDbFlag(m)

		if err == nil {
			err = m.handler.RunForce()
		}
	case "version":
		err = validateDbFlag(m)

		if err == nil {
			err = m.handler.RunVersion()
		}
	default:
		log.Fatalf("Invalid command: %s. Use 'up', 'down', 'force', 'version', or 'rename'", m.handler.MigrationConfig.Command)
	}

	if err != nil {
		log.Fatal(err)
	}

}

func validateDbFlag(m *MigrationCLI) (err error) {

	if m.handler.MigrationConfig.DatabaseURL == "" {

		return fmt.Errorf("DATABASE_URL must be set via the -db flag %v", err)
	}

	return err
}
