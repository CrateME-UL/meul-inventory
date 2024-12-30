package postgres_migrations

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

func NewMigrationFilesHandler(migrationFilesOrderHandler *MigrationFilesOrderHandler) (migrationFilesHandler *MigrationFilesHandler) {

	return &MigrationFilesHandler{
		MigrationFilesOrderHandler: migrationFilesOrderHandler,
	}
}

type MigrationFilesHandler struct {
	MigrationFilesOrderHandler *MigrationFilesOrderHandler
}

func (m *MigrationFilesHandler) renameMigrationFiles(baseName string) (err error) {
	dir := filepath.Dir(baseName)
	name := filepath.Base(baseName)
	timestamp := time.Now().Format("20060102150405")
	order, err := m.MigrationFilesOrderHandler.getNextMigrationOrder(dir)

	if err != nil {

		return fmt.Errorf("failed to get next migration order: %v", err)
	}

	orderPrefix := fmt.Sprintf("%04d", order)
	upFile := filepath.Join(dir, fmt.Sprintf("%s.up.sql", name))
	downFile := filepath.Join(dir, fmt.Sprintf("%s.down.sql", name))
	newUpFile := filepath.Join(dir, fmt.Sprintf("%s_%s_%s.up.sql", orderPrefix, timestamp, name))
	newDownFile := filepath.Join(dir, fmt.Sprintf("%s_%s_%s.down.sql", orderPrefix, timestamp, name))
	fmt.Printf("Looking for files:\n - %s\n - %s\n", upFile, downFile)
	fmt.Printf("Renaming to:\n - %s\n - %s\n", newUpFile, newDownFile)

	if err := os.Rename(upFile, newUpFile); err != nil {

		return fmt.Errorf("failed to rename .up.sql file: %w", err)
	}

	if err := os.Rename(downFile, newDownFile); err != nil {

		return fmt.Errorf("failed to rename .down.sql file: %w", err)
	}

	fmt.Printf("Files renamed successfully:\n")
	fmt.Printf("Old .up.sql: %s -> New .up.sql: %s\n", upFile, newUpFile)
	fmt.Printf("Old .down.sql: %s -> New .down.sql: %s\n", downFile, newDownFile)

	return err
}

type MigrationFilesOrderHandler struct{}

func NewMigrationFilesOrderHandler() (migrationFilesOrderHandler *MigrationFilesOrderHandler) {

	return &MigrationFilesOrderHandler{}
}

func (m *MigrationFilesOrderHandler) getNextMigrationOrder(dir string) (order int, err error) {
	files, err := os.ReadDir(dir)
	var highestOrder int = -1

	if err != nil {

		return highestOrder, fmt.Errorf("failed to read migration directory: %w", err)
	}

	for _, file := range files {

		if file.IsDir() {

			continue
		}

		matches := regexp.MustCompile(`^(\d{4})_\d{14}_.+\.up\.sql$`).FindStringSubmatch(file.Name())

		if len(matches) > 0 {
			order, err := strconv.Atoi(matches[1])

			if err != nil {

				return highestOrder, fmt.Errorf("invalid migration order in file %s: %w", file.Name(), err)
			}

			if order > highestOrder {
				highestOrder = order
			}
		}
	}

	return highestOrder + 1, err
}
