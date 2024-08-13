package pkg

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// readsfrom a specific line from a file(banner)
func GetLine(filename string, num int) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1 // Line numbers start from 1
	for scanner.Scan() {
		if lineNum == num {
			return scanner.Text()
		}
		lineNum++
	}

	return ""
}

func GenerateASCII(input, banner string) string {
	var result strings.Builder

	// Iterate over each line (8 lines per character)
	for line := 1; line < 9; line++ {
		// Iterate over each character in input
		for _, char := range input {
			pos := 1 + int(char-' ')*9 + line
			lineContent := GetLine(banner, pos)
			result.WriteString(lineContent)
		}
		result.WriteString("\n") // New line after each set of characters
	}

	return result.String()
}

func WriteToOutputFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", filename, err)
	}
	defer file.Close()

	_, err = fmt.Fprint(file, content)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %v", filename, err)
	}

	return nil
}

func ValidateFlagFormat() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}

	re := regexp.MustCompile(`^--\w+=.*`)
	arg := os.Args[1]

	if !re.MatchString(arg) {
		fmt.Println("usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	// return nil
}

func ValidateOutputFileName(outputFileName string) error {
	bannedFiles := map[string]bool{
		"standard.txt":   true,
		"shadow.txt":     true,
		"thinkertoy.txt": true,
	}

	if outputFileName == "" {
		return fmt.Errorf("error: output file name is required")
	} else if bannedFiles[outputFileName] {
		return fmt.Errorf("error: %v is a banner file, use another file name", outputFileName)
	} else if strings.HasSuffix(outputFileName, "/standard.txt") ||
		strings.HasSuffix(outputFileName, "/shadow.txt") ||
		strings.HasSuffix(outputFileName, "/thinkertoy.txt") {
		return fmt.Errorf("error: direct access to the default banner file path is not allowed. '%v'", outputFileName)
	}

	return nil
}
