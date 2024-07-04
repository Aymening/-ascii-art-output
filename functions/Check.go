package handler

import (
	"fmt"
	"os"
)

func Check(str string) {
	// Check if string contains valid characters
	for i := 0; i < len(str); i++ {
		if str[i] < 32 || str[i] > 126 {
			fmt.Println("PROGRAM : Invalid characters")
			os.Exit(0)
		}
	}
}
