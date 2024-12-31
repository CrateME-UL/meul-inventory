package postgres_migrations

import (
	"reflect"
	"testing"
)

func TestNewMigrationFilesHandler(t *testing.T) {
	type args struct {
		migrationFilesOrderHandler *MigrationFilesOrderHandler
	}
	tests := []struct {
		name                      string
		args                      args
		wantMigrationFilesHandler *MigrationFilesHandler
	}{
		{
			name: "whenNewMigrationFilesHandler_thenReturnsMigrationFilesHandler",
			args: args{
				migrationFilesOrderHandler: &MigrationFilesOrderHandler{},
			},
			wantMigrationFilesHandler: &MigrationFilesHandler{
				MigrationFilesOrderHandler: &MigrationFilesOrderHandler{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMigrationFilesHandler := DefaultMigrationFilesHandler(tt.args.migrationFilesOrderHandler); !reflect.DeepEqual(gotMigrationFilesHandler, tt.wantMigrationFilesHandler) {
				t.Errorf("NewMigrationFilesHandler() = %v, want %v", gotMigrationFilesHandler, tt.wantMigrationFilesHandler)
			}
		})
	}
}

// todo: add fs implementation
// func TestMigrationFilesHandler_RenameMigrationFiles(t *testing.T) {
// 	type fields struct {
// 		MigrationFilesOrderHandler *MigrationFilesOrderHandler
// 	}
// 	type args struct {
// 		baseName string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "rename migration files successfully",
// 			fields: fields{
// 				MigrationFilesOrderHandler: &MigrationFilesOrderHandler{}, // Mock handler if needed
// 			},
// 			args: args{
// 				baseName: "migration_1",
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "error when renaming a non-existing file",
// 			fields: fields{
// 				MigrationFilesOrderHandler: &MigrationFilesOrderHandler{},
// 			},
// 			args: args{
// 				baseName: "non_existing_file",
// 			},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Set up the mock filesystem
// 			fs := afero.NewMemMapFs()
// 			// Create some initial files in the filesystem
// 			_ = afero.WriteFile(fs, "/testdir/migration_1.sql", []byte("content1"), 0644)
// 			_ = afero.WriteFile(fs, "/testdir/migration_2.sql", []byte("content2"), 0644)

// 			// Create a handler with the mock filesystem
// 			m := &MigrationFilesHandler{
// 				MigrationFilesOrderHandler: tt.fields.MigrationFilesOrderHandler,
// 				// Injecting the filesystem here (assuming MigrationFilesOrderHandler interacts with fs)
// 				FS: fs,
// 			}

// 			// Call RenameMigrationFiles with the test arguments
// 			err := m.RenameMigrationFiles(tt.args.baseName)

// 			// Check for errors and assert the correct behavior
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("MigrationFilesHandler.RenameMigrationFiles() error = %v, wantErr %v", err, tt.wantErr)
// 			}

// 			// Verify file renaming or absence of it in the filesystem based on the test
// 			_, err = fs.Stat("/testdir/" + tt.args.baseName + "_renamed.sql")
// 			if tt.wantErr && err == nil {
// 				t.Errorf("Expected error when renaming but file was found")
// 			} else if !tt.wantErr && err != nil {
// 				t.Errorf("Expected no error when renaming but got: %v", err)
// 			}
// 		})
// 	}
// }

func TestNewMigrationFilesOrderHandler(t *testing.T) {
	tests := []struct {
		name                           string
		wantMigrationFilesOrderHandler *MigrationFilesOrderHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMigrationFilesOrderHandler := DefaultMigrationFilesOrderHandler(); !reflect.DeepEqual(gotMigrationFilesOrderHandler, tt.wantMigrationFilesOrderHandler) {
				t.Errorf("NewMigrationFilesOrderHandler() = %v, want %v", gotMigrationFilesOrderHandler, tt.wantMigrationFilesOrderHandler)
			}
		})
	}
}

func TestMigrationFilesOrderHandler_GetNextMigrationOrder(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name      string
		m         *MigrationFilesOrderHandler
		args      args
		wantOrder int
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MigrationFilesOrderHandler{}
			gotOrder, err := m.GetNextMigrationOrder(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("MigrationFilesOrderHandler.GetNextMigrationOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOrder != tt.wantOrder {
				t.Errorf("MigrationFilesOrderHandler.GetNextMigrationOrder() = %v, want %v", gotOrder, tt.wantOrder)
			}
		})
	}
}
