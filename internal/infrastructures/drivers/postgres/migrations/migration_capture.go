package migrations

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func (m *MigrationHandler) CatchMigrationsToSQLFiles() ([]string, error) {
	inputFile := "../internal/infrastructures/drivers/postgres/migrations/log/migration.log"
	outputDir := "../internal/infrastructures/drivers/postgres/migrations/sql"
	var outputFiles []string

	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current Directory:", dir)

	// Create the output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		fmt.Println("Error creating output directory:", err)
		return nil, err
	}

	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return nil, err
	}
	defer file.Close()

	// Regular expression to match CREATE, ALTER, DELETE statements
	re := regexp.MustCompile(`(?i)\b(CREATE|ALTER|DELETE)\b.*`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			// Generate a filename based on the first few words of the SQL command
			words := strings.Fields(line)
			filename := strings.ToLower(strings.Join(words[2:5], "_"))
			filename = strings.ReplaceAll(filename, `"`, "")
			filename = strings.ReplaceAll(filename, ";", "")

			// Remove the first two words from the line
			if len(words) > 2 {
				line = strings.Join(words[2:], " ")
			}

			// Write the line to the new file
			outputFileUp := filepath.Join(outputDir, filename+".up.sql")
			outputFileDown := filepath.Join(outputDir, filename+".down.sql")
			if err := os.WriteFile(outputFileUp, []byte(line), 0644); err != nil {
				fmt.Println("Error writing to output file:", err)
				return nil, err
			}

			//write another file for down migration that is empty
			var rollbackBuilder strings.Builder
			rollbackBuilder.WriteString("-- TODO: Add rollback migration here.\n")
			rollbackBuilder.WriteString("DO $$\n")
			rollbackBuilder.WriteString("BEGIN\n")
			rollbackBuilder.WriteString("    RAISE EXCEPTION 'TODO: Rollback migration not implemented yet!';\n")
			rollbackBuilder.WriteString("END;\n")
			rollbackBuilder.WriteString("$$;")
			rollbackContent := rollbackBuilder.String()
			if err := os.WriteFile(outputFileDown, []byte(rollbackContent), 0644); err != nil {
				fmt.Println("Error writing to output file:", err)
				return nil, err
			}

			outputFiles = append(outputFiles, strings.TrimSuffix(outputFileUp, ".up.sql"))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
	}

	return outputFiles, nil
}
