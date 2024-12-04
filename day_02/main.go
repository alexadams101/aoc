package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Invalid number of args passed")
		return
	}
	
	data := os.Args[0]
}