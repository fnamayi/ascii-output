package pkg

import (
	"fmt"
	"os"
	"strings"
)

// printUsage prints the usage message
func PrintUsage() {
	fmt.Println("Usage: go run . --output=<fileName.txt> [STRING] [BANNER]")
	fmt.Println("EX: go run . --output=<fileName.txt> something standard")
}

// getASCIIContent reads the appropriate ASCII art file based on the banner argument
func GetASCIIContent(banner string) ([]string, error) {
	var fileName string
	switch banner {
	case "standard":
		fileName = "standard.txt"
	case "shadow":
		fileName = "shadow.txt"
	case "thinkertoy":
		fileName = "thinkertoy.txt"
	default:
		return nil, fmt.Errorf("unsupported banner: %s", banner)
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(content),"\n"), nil
}

// generateASCIIArt generates ASCII art for the given input string using the ASCII content
func GenerateASCIIArt(input string, asciiContent []string) string {
	var result strings.Builder
	for _, char := range input {
		if char >= 32 && char <= 126 {
			start := (int(char) - 32) * 9
			for i := 0; i < 8; i++ {
				result.WriteString(asciiContent[start+i])
				result.WriteString("\n")
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}