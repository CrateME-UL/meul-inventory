package migrations_resource

import (
	"flag"
	"os"
)

// MigrationFlags holds all the command-line flag values
type MigrationFlags struct {
	DatabaseURL    string
	MigrationsPath string
	Command        string
	Steps          int
	ForceVersion   int
	BaseName       string
}

// NewMigrationFlags parses and returns MigrationFlags
func NewMigrationFlags() MigrationFlags {
	flags := MigrationFlags{}
	flag.StringVar(&flags.DatabaseURL, "db", os.Getenv("DATABASE_URL"), "Database connection URL")
	flag.StringVar(&flags.MigrationsPath, "path", "file://migrations", "Path to migrations directory")
	flag.StringVar(&flags.Command, "command", "up", "Migration command: up, down, force, version, or rename")
	flag.IntVar(&flags.Steps, "steps", 1, "Number of steps for the 'down' command")
	flag.IntVar(&flags.ForceVersion, "version", 0, "Version to force with the 'force' command")
	flag.StringVar(&flags.BaseName, "base", "", "Base name of the migration file (without '.up.sql' or '.down.sql')")
	flag.Parse()
	return flags
}
