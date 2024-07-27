package main

import (
    "flag"
    "os"
    "testing"
)

// TestFlagParsing tests the flag parsing functionality
func TestFlagParsing(t *testing.T) {
    tests := []struct {
        args                  []string
        expectedOutputFileName string
    }{
        {
            args:                  []string{"program", "-output=filename.txt"},
            expectedOutputFileName: "filename.txt",
        },
        {
            args:                  []string{"program", "-output="},
            expectedOutputFileName: "",
        },
        {
            args:                  []string{"program"},
            expectedOutputFileName: "",
        },
    }

    for _, tt := range tests {
        // Create a new FlagSet for each test case to avoid conflicts
        fs := flag.NewFlagSet("test", flag.ContinueOnError)
        var outputFileName string
        fs.StringVar(&outputFileName, "output", "", "Output file name (required)")

        // Set the command-line arguments
        os.Args = tt.args

        // Parse the command-line arguments
        if err := fs.Parse(tt.args[1:]); err != nil {
            t.Fatalf("Failed to parse flags: %v", err)
        }

        // Check the result
        if outputFileName != tt.expectedOutputFileName {
            t.Errorf("For args %v, expected outputFileName %q but got %q", tt.args, tt.expectedOutputFileName, outputFileName)
        }
    }
}
