package main

import "fmt"

var arr = []int{245, 154, 568, 324, 654, 325}

// InsertionSort sorts an array using the insertion sort algorithm.
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Move elements of arr[0..i-1], that are greater than key,
		// to one position ahead of their current position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
		//fmt.Println("key", key, "- Step", i, ":", arr)
	}
	return arr
}

// main function to test the InsertionSort function
func main() {
	fmt.Println("- Original array :", arr)
	sortedArr := InsertionSort(arr)
	fmt.Println("- Sorted array :", sortedArr)
}
