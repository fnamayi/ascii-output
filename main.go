package main

import (
	"fmt"
	"os"

	ascii "ascii-output/pkg"
)

func main() {
	input, _ := ascii.InputArgs(os.Args)
	asciiFields := ascii.FileChoice(os.Args)
	printWords(input, asciiFields)
}

// printing of the characters
func printWords(input []string, asciiFields []string) {
	for _, word := range input {
		if word == "" {
			fmt.Println()
		} else {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if !validChar(char) {
						return
					}
					startPoint := Start(int(char))
					fmt.Print(asciiFields[startPoint+i])
				}
				fmt.Println()
			}
		}
	}
}

// starting positions of the characters
func Start(s int) int {
	pos := int(s-32)*9 + 1
	return pos
}

func validChar(s rune) bool {
	if !(s >= ' ' && s <= '~') {
		fmt.Println("Error:" + string(s) + " " + "is not valid character input")
		return false
	}
	return true
}
