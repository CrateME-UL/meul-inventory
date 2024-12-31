package postgres_migrations

import (
	"reflect"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func TestMigrationHandler_initializeMigration(t *testing.T) {
	type fields struct {
		DatabaseURL           string
		MigrationsPath        string
		Command               string
		Steps                 int
		ForceVersion          int
		BaseName              string
		MigrationFilesHandler *MigrationFilesHandler
	}
	tests := []struct {
		name    string
		fields  fields
		want    *migrate.Migrate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MigrationHandler{
				DatabaseURL:           tt.fields.DatabaseURL,
				MigrationsPath:        tt.fields.MigrationsPath,
				Command:               tt.fields.Command,
				Steps:                 tt.fields.Steps,
				ForceVersion:          tt.fields.ForceVersion,
				BaseName:              tt.fields.BaseName,
				MigrationFilesHandler: tt.fields.MigrationFilesHandler,
			}
			got, err := m.initializeMigration()
			if (err != nil) != tt.wantErr {
				t.Errorf("MigrationHandler.initializeMigration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MigrationHandler.initializeMigration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMigrationHandler_executeMigration(t *testing.T) {
	type fields struct {
		DatabaseURL           string
		MigrationsPath        string
		Command               string
		Steps                 int
		ForceVersion          int
		BaseName              string
		MigrationFilesHandler *MigrationFilesHandler
	}
	type args struct {
		action func(*migrate.Migrate) error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MigrationHandler{
				DatabaseURL:           tt.fields.DatabaseURL,
				MigrationsPath:        tt.fields.MigrationsPath,
				Command:               tt.fields.Command,
				Steps:                 tt.fields.Steps,
				ForceVersion:          tt.fields.ForceVersion,
				BaseName:              tt.fields.BaseName,
				MigrationFilesHandler: tt.fields.MigrationFilesHandler,
			}
			if err := m.executeMigration(tt.args.action); (err != nil) != tt.wantErr {
				t.Errorf("MigrationHandler.executeMigration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMigrationHandler_RunUp(t *testing.T) {
	type fields struct {
		DatabaseURL           string
		MigrationsPath        string
		Command               string
		Steps                 int
		ForceVersion          int
		BaseName              string
		MigrationFilesHandler *MigrationFilesHandler
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MigrationHandler{
				DatabaseURL:           tt.fields.DatabaseURL,
				MigrationsPath:        tt.fields.MigrationsPath,
				Command:               tt.fields.Command,
				Steps:                 tt.fields.Steps,
				ForceVersion:          tt.fields.ForceVersion,
				BaseName:              tt.fields.BaseName,
				MigrationFilesHandler: tt.fields.MigrationFilesHandler,
			}
			if err := m.RunUp(); (err != nil) != tt.wantErr {
				t.Errorf("MigrationHandler.RunUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMigrationHandler_RunDown(t *testing.T) {
	type fields struct {
		DatabaseURL           string
		MigrationsPath        string
		Command               string
		Steps                 int
		ForceVersion          int
		BaseName              string
		MigrationFilesHandler *MigrationFilesHandler
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MigrationHandler{
				DatabaseURL:           tt.fields.DatabaseURL,
				MigrationsPath:        tt.fields.MigrationsPath,
				Command:               tt.fields.Command,
				Steps:                 tt.fields.Steps,
				ForceVersion:          tt.fields.ForceVersion,
				BaseName:              tt.fields.BaseName,
				MigrationFilesHandler: tt.fields.MigrationFilesHandler,
			}
			if err := m.RunStepsDown(); (err != nil) != tt.wantErr {
				t.Errorf("MigrationHandler.RunDown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMigrationHandler_RunRename(t *testing.T) {
	type fields struct {
		DatabaseURL           string
		MigrationsPath        string
		Command               string
		Steps                 int
		ForceVersion          int
		BaseName              string
		MigrationFilesHandler *MigrationFilesHandler
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MigrationHandler{
				DatabaseURL:           tt.fields.DatabaseURL,
				MigrationsPath:        tt.fields.MigrationsPath,
				Command:               tt.fields.Command,
				Steps:                 tt.fields.Steps,
				ForceVersion:          tt.fields.ForceVersion,
				BaseName:              tt.fields.BaseName,
				MigrationFilesHandler: tt.fields.MigrationFilesHandler,
			}
			if err := m.RunRename(); (err != nil) != tt.wantErr {
				t.Errorf("MigrationHandler.RunRename() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMigrationHandler_RunForce(t *testing.T) {
	type fields struct {
		DatabaseURL           string
		MigrationsPath        string
		Command               string
		Steps                 int
		ForceVersion          int
		BaseName              string
		MigrationFilesHandler *MigrationFilesHandler
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MigrationHandler{
				DatabaseURL:           tt.fields.DatabaseURL,
				MigrationsPath:        tt.fields.MigrationsPath,
				Command:               tt.fields.Command,
				Steps:                 tt.fields.Steps,
				ForceVersion:          tt.fields.ForceVersion,
				BaseName:              tt.fields.BaseName,
				MigrationFilesHandler: tt.fields.MigrationFilesHandler,
			}
			if err := m.RunForce(); (err != nil) != tt.wantErr {
				t.Errorf("MigrationHandler.RunForce() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMigrationHandler_RunVersion(t *testing.T) {
	type fields struct {
		DatabaseURL           string
		MigrationsPath        string
		Command               string
		Steps                 int
		ForceVersion          int
		BaseName              string
		MigrationFilesHandler *MigrationFilesHandler
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MigrationHandler{
				DatabaseURL:           tt.fields.DatabaseURL,
				MigrationsPath:        tt.fields.MigrationsPath,
				Command:               tt.fields.Command,
				Steps:                 tt.fields.Steps,
				ForceVersion:          tt.fields.ForceVersion,
				BaseName:              tt.fields.BaseName,
				MigrationFilesHandler: tt.fields.MigrationFilesHandler,
			}
			if err := m.RunVersion(); (err != nil) != tt.wantErr {
				t.Errorf("MigrationHandler.RunVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
