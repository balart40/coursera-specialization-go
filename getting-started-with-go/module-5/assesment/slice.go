package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop,
the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/
func main() {
	// Create slice size 3
	var sli = make([]int, 0, 3)
	var str string
	// do infinite loop
	for i := 0; ; i++ {
		fmt.Print("Provide an integer, x to exit: ")
		// scan for input
		_, err := fmt.Scan(&str)
		// if error fail
		if err != nil {
			fmt.Println("Error scanning input")
		}
		// check first if want to exit if so break
		if str == "x" || str == "X" {
			break
		}
		// transform to int
		integer, err := strconv.Atoi(str)
		// if error fail
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		} else {
			sli = append(sli, integer)
			// make a copy
			copySli := make([]int, len(sli))
			copy(copySli, sli)
			sort.Ints(copySli)
			fmt.Println("Sorted: ", copySli)
		}
	}
}
