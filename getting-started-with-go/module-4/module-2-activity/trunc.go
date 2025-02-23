package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Printf("Provide a floating poinit number!: ")
	// Scan for floating point capture error if any
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println("Error reading input")
	}
	// get the dot index
	dotIndex := strings.Index(s, ".")
	// no decimal point return as it is
	if dotIndex == -1 {
		fmt.Printf(s)
	}
	if dotIndex == 0 {
		fmt.Printf("0")
	}
	// if we get to this point just return everything without the point
	fmt.Printf("\nNumber Truncated is: %s\n", s[0:dotIndex])
}
