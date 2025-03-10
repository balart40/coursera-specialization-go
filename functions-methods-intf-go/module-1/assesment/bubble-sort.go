package main

import "fmt"

/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.

Submit your Go program source code.
*/
func Swap(sli []int, i int) {
	sli[i], sli[i+1] = sli[i+1], sli[i]
}

func BubbleSort(sli []int) {
	n := len(sli)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func main() {
	// *************  request the 10 nums *************
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Printf("Provide number #%d:", i+1)
		_, err := fmt.Scan(&arr[i]) // Read input and store it in the array
		if err != nil {
			fmt.Println("Error scanning name input")
		}
	}
	// *************  call bubble sort *************
	BubbleSort(arr)
	fmt.Println(arr)
}
