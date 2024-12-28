package postgres_migrations

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

// todo: add a struct that hold these values to be able to test them
func renameMigrationFiles(baseName string) {
	dir := filepath.Dir(baseName)
	name := filepath.Base(baseName)

	timestamp := time.Now().Format("20060102150405")
	orderPrefix := fmt.Sprintf("%04d", getNextMigrationOrder(dir))

	upFile := filepath.Join(dir, fmt.Sprintf("%s.up.sql", name))
	downFile := filepath.Join(dir, fmt.Sprintf("%s.down.sql", name))

	newUpFile := filepath.Join(dir, fmt.Sprintf("%s_%s_%s.up.sql", orderPrefix, timestamp, name))
	newDownFile := filepath.Join(dir, fmt.Sprintf("%s_%s_%s.down.sql", orderPrefix, timestamp, name))

	fmt.Printf("Looking for files:\n - %s\n - %s\n", upFile, downFile)
	fmt.Printf("Renaming to:\n - %s\n - %s\n", newUpFile, newDownFile)

	if err := os.Rename(upFile, newUpFile); err != nil {
		log.Fatalf("Failed to rename .up.sql file: %v", err)
	}
	if err := os.Rename(downFile, newDownFile); err != nil {
		log.Fatalf("Failed to rename .down.sql file: %v", err)
	}

	fmt.Printf("Files renamed successfully:\n")
	fmt.Printf("Old .up.sql: %s -> New .up.sql: %s\n", upFile, newUpFile)
	fmt.Printf("Old .down.sql: %s -> New .down.sql: %s\n", downFile, newDownFile)
}

// todo: add a struct that hold these values to be able to test them
func getNextMigrationOrder(dir string) int {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("Failed to read migration directory: %v", err)
	}

	var highestOrder int = -1
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		matches := regexp.MustCompile(`^(\d{4})_\d{14}_.+\.up\.sql$`).FindStringSubmatch(file.Name())
		if len(matches) > 0 {
			order, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Printf("Invalid migration order in file %s: %v", file.Name(), err)
				continue
			}

			if order > highestOrder {
				highestOrder = order
			}
		}
	}

	return highestOrder + 1
}
