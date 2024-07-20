package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	ascii "ascii-output/pkg"
)

func main() {
	if len(os.Args) <= 1 {
		Errormsg()
		os.Exit(0)
		// return

	}
	// Declare flag
	var outputFileName string
	flag.StringVar(&outputFileName, "output", "", "Output file name (required)")

	flag.Parse()

	if outputFileName == "" {
		Errormsg()
		os.Exit(1)
	} else if outputFileName == "standard.txt" || outputFileName == "shadow.txt" || outputFileName == "thinkertoy.txt" {
		fmt.Println("Error: You can not write on these banner files.")
		return
	} else if strings.HasSuffix(outputFileName, "/standard.txt") || strings.HasSuffix(outputFileName, "/shadow.txt") || strings.HasSuffix(outputFileName, "/thinkertoy.txt") {
		fmt.Println("Warning: Attempt to edit banner f ile")
		return
	}

	// Remaining arguments after flags
	args := flag.Args()

	var banner, input string

	if len(args) == 1 {
		input = args[0]
		banner = "standard"
	} else if len(args) == 2 {
		input = args[0]
		banner = strings.ToLower(args[1])
	} else {
		Errormsg()
		os.Exit(1)
	}

	switch banner {
	case "shadow":
		banner = "shadow.txt"
	case "thinkertoy":
		banner = "thinkertoy.txt"

	case "standard":
		banner = "standard.txt"
	}

	// Generate ASCII art
	result := ascii.GenerateASCII(input, banner)

	// Write result to file
	err := ascii.WriteToOutputFile(outputFileName, result)
	if err != nil {
		fmt.Println("Error Writting to a file %:", err)
		os.Exit(1)
	}
}

func Errormsg() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
}
