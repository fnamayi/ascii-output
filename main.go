package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ascii-output/pkg"
)

func main() {
	// if len(os.Args) > 2 {

	// }

	pkg.ValidateFlagFormat()

	// Define and parse flags
	var outputFileName string
	flag.StringVar(&outputFileName, "output", "", "Output file name (optional)")

	// Parse command-line flags
	flag.Parse()

	// Extract arguments after flags
	args := flag.Args()

	if len(args) < 1 {
		printUsage()
		os.Exit(0)
	}

	// Extract input and optional banner from arguments
	input := args[0]
	var banner string

	if len(args) > 1 {
		banner = strings.ToLower(args[1])
	} else {
		banner = "standard" // Default banner if none is provided
	}

	// Map banner to file name
	switch banner {
	case "shadow":
		banner = "shadow.txt"
	case "thinkertoy":
		banner = "thinkertoy.txt"
	case "standard":
		banner = "standard.txt"
	default:
		fmt.Printf("Unknown banner type: %s\n", banner)
		Errormsg()
		os.Exit(1)
	}

	// Generate ASCII art
	result := pkg.GenerateASCII(input, banner)

	// Check if output file name is provided
	if outputFileName != "" {
		// Validate the output file name
		if err := pkg.ValidateOutputFileName(outputFileName); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Write result to file
		if err := pkg.WriteToOutputFile(outputFileName, result); err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}
	} else {
		// Print result to terminal
		fmt.Println(result)
	}
}

func Errormsg() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
}
