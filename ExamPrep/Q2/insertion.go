package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	// loop through whole array starting on the second item
	for i := 1; i < n; i++ {
		// save value of current index in tmp var
		item := arr[i]
		// set start point for backwards loop to the previous element in the list
		j := i - 1
		// loop backwards until the start of the array
		// or until value of the tmp var is not less
		for j >= 0 && item < arr[j] {
			// shfit the element to the right (one up)
			arr[j+1] = arr[j]
			j--
		}
		// set the tmp value to its new correct position
		arr[j+1] = item
		fmt.Println(arr)
	}
}
