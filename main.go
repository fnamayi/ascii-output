package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	ascii "ascii-output/pkg"
)

func main() {
	ascii.ValidateFlagFormat()

	if len(os.Args) <= 1 {
		Errormsg()
		os.Exit(0)
		// return

	}
	// Declare flag
	var outputFileName string
	flag.StringVar(&outputFileName, "output", "", "Output file name (required)")

	flag.Parse()

	if err := ascii.ValidateOutputFileName(outputFileName); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Remaining arguments after flags
	args := flag.Args()

	var banner, input string

	if len(args) == 1 {
		input = args[0]
		// input = strings.ReplaceAll(input, "/n", "//n")

		banner = "standard"
	} else if len(args) == 2 {
		input = args[0]
		// input = strings.ReplaceAll(input, "/n", "//n")

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


	// Generate ASCII art and store in variable result
	result := ascii.GenerateASCII(input, banner)

	// Write result to file after error checking
	err := ascii.WriteToOutputFile(outputFileName, result)
	if err != nil {
		fmt.Println("Error Writting to a file %:", err)
		os.Exit(1)
	}
}

func Errormsg() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
}
