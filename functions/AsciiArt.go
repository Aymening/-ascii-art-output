package handler

import (
	"fmt"
	"strings"
)

func AsciiArt(arg string, cont string) {
	// Split banner content by double newline
	SplitFile := strings.Split(cont, "\n\n")
	// Split input argument by "\n" escape sequence
	SplitArg := strings.Split(arg, "\\n")
	// Initialize 2D slice to store result
	result := make([][]string, len(SplitArg))
	for mak := 0; mak < len(SplitArg); mak++ {
		result[mak] = make([]string, 8)
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
				// Print each line of ASCII art
				fmt.Println(result[i][l])
			} else {
				fmt.Println()
				break
			}
		}
	}
}

