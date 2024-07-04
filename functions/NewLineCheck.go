package handler

import "strings"

func NewLineCheck(str string) bool {
	// Replace "\\n" with "\n" in the string
	result := strings.ReplaceAll(str, "\\n", "\n")
	cont := 0
	// Count newline characters
	for i := 0; i < len(result); i++ {
		if result[i] == '\n' {
			cont++
		}
	}
	// Check if number of newline characters equals length of string
	return cont == len(result)
}
