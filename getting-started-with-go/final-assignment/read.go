package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs.

Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct
which contains the first and last names found in the file.

Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file.

After reading all lines from the file,
your program should iterate through your slice of structs and print the first and last names found in each struct.

Submit your source code for the program, “read.go”.
*/

type Name struct {
	fname string
	lname string
}

func getFirst20Characters(str string) string {
	if len(str) > 20 {
		strToRunes := []rune(str)
		return string(strToRunes[0:20])
	} else {
		return str
	}
}

func get_file_name() string {
	var fileName string
	fmt.Println("Provide a fileName:")
	_, err := fmt.Scan(&fileName)
	if err != nil {
		fmt.Println("Error scanning fileName input")
	}
	return fileName
}

func process_file(fileName string) []Name {
	var people []Name
	var nameObject Name
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	// Ensure the file is closed when the function completes.
	defer file.Close()
	// Create a new scanner for the file.
	scanner := bufio.NewScanner(file)

	// Loop through the file line by line.
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), " ")
		names[0] = getFirst20Characters(names[0])
		names[1] = getFirst20Characters(names[1])
		nameObject.fname, nameObject.lname = names[0], names[1]
		people = append(people, nameObject)
	}
	return people
}

func main() {
	fileName := get_file_name()
	peopleSlice := process_file(fileName)
	for _, nameStruct := range peopleSlice {
		fmt.Println(nameStruct.fname, nameStruct.lname)
	}
}
