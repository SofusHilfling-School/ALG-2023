package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i // Return the index of the found element
		}
		fmt.Println("index:", i)
	}
	return -1 // Return -1 if the element is not found
}

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		fmt.Printf("index: %d, value: %d\n", mid, arr[mid])

		if arr[mid] == target {
			return mid // Return the index of the found element
		} else if arr[mid] < target {
			low = mid + 1 // Continue searching in the right half
		} else {
			high = mid - 1 // Continue searching in the left half
		}
	}

	return -1 // Return -1 if the element is not found
}
