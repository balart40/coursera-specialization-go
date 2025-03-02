package main

import (
	"fmt"
	"strings"
)

/*
Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’,
and contains the character ‘a’. The program should print “Not Found!” otherwise.
The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
*/
func main() {
	var str string
	var prefix, suffix string = "i", "n"
	var inner = "a"

	fmt.Println("Provide a string:")

	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println("Error scanning input")
	}
	strLower := strings.ToLower(str)

	hasPrefix := strings.HasPrefix(strLower, prefix)
	hasSuffix := strings.HasSuffix(strLower, suffix)
	hasInner := strings.Contains(strLower, inner)

	if hasPrefix && hasSuffix && hasInner {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
