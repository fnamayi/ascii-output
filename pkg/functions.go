package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
