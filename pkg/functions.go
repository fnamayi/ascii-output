package pkg

import (
	"fmt"
	"os"
	"strings"
)

// InputArgs function processes command-line arguments
func InputArgs(osArgs []string) ([]string, int) {
	var args string
	if !(len(osArgs) == 2 || len(osArgs) == 3) {
		fmt.Println("usage: go run . <input string> <optional flag>")
		return []string{}, 0
	} else if len(osArgs) == 3 && !(osArgs[2] == "-t" || osArgs[2] == "-sh") {
		fmt.Println("Only one string is accepted")
		return []string{}, 0
	} else {
		args = osArgs[1]
	}

	args = strings.ReplaceAll(args, "\n", "\\n")
	args = strings.ReplaceAll(args, "\\t", " ")

	input := strings.Split(args, "\\n")

	if input[0] == "" && len(input) == 1 {
		os.Exit(0)
	} else if input[0] == "" && input[1] == "" && len(input) == 2 {
		fmt.Println()
		os.Exit(0)
	}
	return input, len(osArgs)
}

// FileChoice function selects the appropriate ASCII art file to read
//changes to be done to accept --output flag
func FileChoice(osArgs []string) []string {
	var file []byte
	var asciiFields []string

	if len(osArgs) == 3 && osArgs[2] == "-t" {
		filec, err := os.ReadFile("thinkertoy.txt")
		file = filec
		asciiField := strings.Split(string(file), "\r\n")
		asciiFields = asciiField
		if Tamper(asciiFields) {
			os.Exit(0)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	} else if len(osArgs) == 3 && osArgs[2] == "-sh" {
		filec, err := os.ReadFile("shadow.txt")
		file = filec
		asciiField := strings.Split(string(file), "\n")
		asciiFields = asciiField
		if Tamper(asciiFields) {
			return []string{}
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	} else if len(osArgs) == 2 {
		filec, err := os.ReadFile("standard.txt")
		file = filec
		asciiField := strings.Split(string(file), "\n")
		asciiFields = asciiField
		if Tamper(asciiFields) {
			os.Exit(0)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}

	return asciiFields
}

func Tamper(s []string) bool {
	if len(s) != 856 {
		fmt.Println("The File Banner used has been tampered with or is Empty")
		return true
	}
	return false
}
