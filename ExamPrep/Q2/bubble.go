package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	// loop through whole array
	for i := 0; i < n-1; i++ {
		// loop through whole array again minus the current number of iterations
		// this is done to bubble the largest element to the end
		for j := 0; j < n-i-1; j++ {
			// check if current item is larger than the next
			if arr[j] > arr[j+1] {
				// bubble largest value up by swapping arr[j] with arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]

				// other method
				// tmp := arr[j+1]
				// arr[j+1] = arr[j]
				// arr[j] = tmp
			}
		}
		fmt.Println(arr)
	}
}

func bubbleSortWithFlag(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Initialize the flag for each pass
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		// If no swaps occurred, the array is already sorted
		if !swapped {
			break
		}
	}
}
