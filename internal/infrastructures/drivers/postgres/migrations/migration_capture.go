package migrations

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func CatchMigrationsToSQLFiles() {
	inputFile := "/home/jaydutemple007/repos/go-exploration/scripts/migration.log"
	outputDir := "/home/jaydutemple007/repos/go-exploration/scripts/sql_files"

	// Create the output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
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
			filename := strings.ToLower(strings.Join(words[:3], "_")) + ".sql"
			filename = strings.ReplaceAll(filename, `"`, "")
			filename = strings.ReplaceAll(filename, ";", "")

			// Write the line to the new file
			outputFile := filepath.Join(outputDir, filename)
			if err := os.WriteFile(outputFile, []byte(line), 0644); err != nil {
				fmt.Println("Error writing to output file:", err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
	}
}
