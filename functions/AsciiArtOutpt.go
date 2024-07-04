package handler

import (
	"os"
	"strings"
)

func AsciiArtOutput(arg string, cont string, File string) {
	// Split banner content by double newline
	SplitFile := strings.Split(cont, "\n\n")
	// Split input argument by "\n" escape sequence
	SplitArg := strings.Split(arg, "\\n")
	// Initialize 2D slice to store result
	result := make([][]string, len(SplitArg))
	for mak := 0; mak < len(SplitArg); mak++ {
		result[mak] = make([]string, 8)
	}
	// Create or open output file
	file, err := os.Create(File)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	// Handle case with new line in input argument
	if NewLineCheck(os.Args[2]) {
		for i := 0; i < len(os.Args[2])/2; i++ {
			file.WriteString(string("\n"))
		}
		return
	}
	// Construct ASCII art based on input argument and banner content
	for i := 0; i < len(SplitArg); i++ {
		for l := 0; l < 8; l++ {
			for j := 0; j < len(SplitArg[i]); j++ {
				// Split characters of banner content
				SplitCharacters := strings.Split(string(SplitFile[SplitArg[i][j]-32]), "\n")
				if SplitCharacters[l] == "" {
					result[i][l] += "      "
				}
				result[i][l] += SplitCharacters[l]
			}
			if result[i][l] != "" {
				file.WriteString(string(result[i][l]+"\n"))
			} else {
				file.WriteString(string("\n"))
				break
			}
		}
	}
}

