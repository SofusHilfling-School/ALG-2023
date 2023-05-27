package main

import "fmt"

func MergeSortNotInplace(arr []int) []int {
	// Base case: if the array has 0 or 1 element, it is already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Divide the array into two halves
	mid := len(arr) / 2
	left := MergeSortNotInplace(arr[:mid])  // Recursive call to sort the left half
	right := MergeSortNotInplace(arr[mid:]) // Recursive call to sort the right half

	// Merge the sorted halves
	return mergeNotInplace(left, right)
}

func mergeNotInplace(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Compare elements from the left and right arrays and merge them in sorted order
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append the remaining elements from the left or right array (if any)
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func MergeSortWithBottomUpInPlace(arr []int) {
	n := len(arr)
	aux := make([]int, len(arr))
	// split into multiple subarrays with a size starting with 1
	// for each iteration increase the size of the subarray by 2 times untill the subarray size is larger than the whole array
	// i.e. merge larger and larger subarrays
	for size := 1; size < n; size *= 2 {
		// find the low and hi for all subarrays
		// since we are merging two sub arrays (lo-mid and mid-hi) we increase lo with the size 2 times each time to merge
		for lo := 0; lo < n-size; lo += size * 2 {
			// find the mid point between the two subarrays
			mid := lo + size - 1
			// find the hi point
			// to not excceed the arrays bound we choose the smaller of either the lengeth of the array
			// or the start of the next subarray minus 1
			hi := min(lo+size*2-1, n-1)
			merge(arr, aux, lo, mid, hi)
			fmt.Println(arr)
		}
	}
}

func MergeSortWithTopDownInPlace(arr []int) {
	aux := make([]int, len(arr))
	sortTopDown(arr, aux, 0, len(arr)-1)
}

func sortTopDown(arr, aux []int, lo, hi int) {
	// Base case: if the array has 0 or 1 element, it is already sorted
	if hi <= lo {
		return
	}

	// Divide the array into two halves
	mid := lo + (hi-lo)/2

	// Recursive call to sort the left half
	sortTopDown(arr, aux, lo, mid)
	// Recursive call to sort the right half
	sortTopDown(arr, aux, mid+1, hi)

	// Merge the sorted halves
	merge(arr, aux, lo, mid, hi)
	fmt.Println(arr)
}

// Merge two sorted arrays into a single sorted array
func merge(arr, aux []int, lo, mid, hi int) {
	// set i to start of the array
	// set j to middle of array
	i, j := lo, mid+1

	// copy the current state of the array to the aux array
	for k := lo; k <= hi; k++ {
		aux[k] = arr[k]
	}

	for k := lo; k <= hi; k++ {
		// check if all elements from the left subarray has been merged
		if i > mid {
			// copy remaining elements from right subarray
			arr[k] = aux[j]
			j++
			// check if all elements from the right subarray has been merged
		} else if j > hi {
			// copy remaining elements from left subarray
			arr[k] = aux[i]
			i++

		} else if aux[i] <= aux[j] {
			// Compare elements from both subarrays and copy the smaller one
			// if i is less than j add i to array
			arr[k] = aux[i]
			i++
		} else {
			// else then j is bigger than i so add j to the array
			arr[k] = aux[j]
			j++
		}
	}
}

// returns the smallest of the two inputs
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
