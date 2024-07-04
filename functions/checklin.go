package handler

import (
	"bufio"
	"fmt"
	"os"
)

func CheckLine(filename string) bool {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file: ", err)
		os.Exit(0)
	}
	// Create a scanner to read the file line by line
	scan := bufio.NewScanner(file)
	con := 0
	// Count the number of lines in the file
	for scan.Scan() {
		con++
	}
	// Check if the number of lines matches the expected count
	return con == 855
}
