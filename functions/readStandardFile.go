package handler

import (
	"fmt"
	"os"
	"strings"
)

func ReadStandardFile(FileName string) string {
	// Construct file path
	file := "Banner/" + FileName + ".txt"
	
	// Check if file content is valid
	if !CheckLine(file) {
		fmt.Println("ERROR: DO NOT TOUCH THE BANNER")
		os.Exit(0)
	}
	
	// Read file content
	cont, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("PROGRAM : ", err)
		os.Exit(0)
	}
	
	// Remove carriage return characters
	modf := strings.NewReplacer(
		"\r", "",
	)
	result := modf.Replace(string(cont))
	return string(result)
}
