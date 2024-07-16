package main

import (
	"flag"
	"fmt"
	"os"
	ascii "ascii-output/pkg"
)

func main() {
	// Define the output flag
	outputFlag := flag.String("output", "", "Specify the output file name as --output=<fileName.txt>")
	// Parse the flags
	flag.Parse()
	args := flag.Args()

	// Check if the correct number of arguments is provided
	if len(args) < 1 || len(args) > 2 || *outputFlag == "" {
		ascii.PrintUsage()
		return
	}

	// Get the STRING argument
	inputString := args[0]

	// Optionally get the BANNER argument
	banner := "standard"
	if len(args) == 2 {
		banner = args[1]
	}

	// Get the ASCII art content based on the banner
	asciiContent, err := ascii.GetASCIIContent(banner)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate the ASCII art for the input string
	result := ascii.GenerateASCIIArt(inputString, asciiContent)

	// Save the result to the specified output file
	err = os.WriteFile(*outputFlag, []byte(result), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Printf("Output saved to %s\n", *outputFlag)

}
