package handler

import (
	"fmt"
	"os"
	"regexp"
)
func HandleInput() {
	// Regular expression to match flag format "--output=<fileName.txt>"
	IsFlag := regexp.MustCompile(`^--output=(([[:print:]])*.txt)$`)
	flag := IsFlag.FindStringSubmatch(os.Args[1])

	// Handling different command-line argument counts
	if len(os.Args) == 4 {
		// If flag format is invalid, print usage instructions
		if flag == nil {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
		Check(os.Args[2])
		File := ReadStandardFile(os.Args[3])
		AsciiArtOutput(os.Args[2], File, flag[1])
	} else if len(os.Args) == 2 {
		// Handling case with only one command argument
		if NewLineCheck(os.Args[1]) {
			for i := 0; i < len(os.Args[1])/2; i++ {
				fmt.Println()
			}
			return
		}
		Check(os.Args[1])
		// If flag format is invalid, use standard file
		if flag == nil {
			File := ReadStandardFile("standard")
			AsciiArt(os.Args[1], File)
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		}

	} else if len(os.Args) == 3 {
		// Handling case with three command arguments
		Check(os.Args[2])
		if NewLineCheck(os.Args[1]) {
			for i := 0; i < len(os.Args[1])/2; i++ {
				fmt.Println()
			}
			return
		}
		if flag == nil {
			File := ReadStandardFile(os.Args[2])
			AsciiArt(os.Args[1], File)
			return
		}
		File := ReadStandardFile("standard")
		AsciiArtOutput(os.Args[2], File, flag[1])
	}
}
