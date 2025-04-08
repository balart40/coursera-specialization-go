package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

/*
Write a program to sort an array of integers. The program should partition the array into 4 parts,
each of which is sorted by a different goroutine. Each partition should be of approximately equal size.

Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete,
the main goroutine should print the entire sorted list.

*/

func ReadValuesFromConsole() (numbers []int, err error) {
	fmt.Println("Please input numbers(separate with space): ")
	br := bufio.NewReader(os.Stdin)
	a, _, err := br.ReadLine() // ReadLine returns
	// the string containing the line,
	// the isPrefix attribute will be false when
	// returning the last fragment of the line.
	// An error occurs if something goes wrong.
	ns := strings.Split(string(a), " ")
	for _, s := range ns {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	return
}
func PartitionArray(numbers []int, k int) (partitions [][]int) {
	n := len(numbers)
	partitions = make([][]int, k)
	lenPartition := n / k

	for j := 0; j < k; j++ {
		start := j * lenPartition
		var end int

		// For the last partition, include all remaining elements
		if j == k-1 {
			end = n
		} else {
			end = (j + 1) * lenPartition
		}

		partitions[j] = numbers[start:end]
		fmt.Println(partitions[j]) // Optional: debug output
	}

	return
}

func Sort(numbers []int, wg *sync.WaitGroup) {
	QuickSort(numbers)

	if wg != nil {
		wg.Done()
	}
}

func QuickSort(a []int) {
	if len(a) < 2 {
		return
	}
	left, right := 0, len(a)-1 // Start from the extrema

	pivot := rand.Int() % len(a) // Choose a random pivot index

	a[pivot], a[right] = a[right], a[pivot] // Put the pivot element
	// in the right place

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuickSort(a[:left])
	QuickSort(a[left+1:])
}

func AppendAndSortPartitions(partitions [][]int) []int {
	var combined []int
	for _, part := range partitions {
		combined = append(combined, part...)
	}
	sort.Ints(combined)
	return combined
}

func main() {
	// number of sub arrays
	var numSubArr = 4
	// waitgroup object
	var wg sync.WaitGroup

	numbers, err := ReadValuesFromConsole()
	if err != nil {
		fmt.Println("there was an error in  ReadValuesFromConsole:", err)
	}

	n := len(numbers)

	// if we dont have at least 5 numbers we dont need to do partitions
	if n <= numSubArr {
		QuickSort(numbers)
		fmt.Printf("Sorted numbers: %v", numbers)
	} else {
		// create partitions
		partitions := PartitionArray(numbers, numSubArr)

		// number of partitions define the wait semaphore
		wg.Add(numSubArr)
		// create a go routing for the number of partitions
		for i := 0; i < numSubArr; i++ {
			go Sort(partitions[i], &wg)
		}
		wg.Wait()
		sorted := AppendAndSortPartitions(partitions)

		fmt.Printf("Sorted array: %v", sorted)
	}
}
