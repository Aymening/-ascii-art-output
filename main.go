package main

import (
	"fmt"
	"os"

	"handler/functions"
)

	func main() {
		if len(os.Args) > 4 ||  len(os.Args) == 1 {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
		handler.HandleInput()
	}
