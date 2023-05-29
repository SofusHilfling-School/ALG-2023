package main

import "fmt"

func main() {
	topDownMergeSortExample()
}

func bubbleSortExample() {
	// Example usage
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Bubble Sort")
	fmt.Println("Original array:", arr)

	bubbleSort(arr)
	fmt.Println("Sorted array:", arr)
	fmt.Println()
}

func insertionSortExample() {
	// Example usage
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Insertion Sort")
	fmt.Println("Original array:", arr)

	insertionSort(arr)
	fmt.Println("Sorted array:", arr)
	fmt.Println()
}

func topDownMergeSortExample() {
	// Example usage
	arr := []int{9, 3, 7, 5, 1, 8, 2, 6, 10, 4}

	fmt.Println("Merge Sort -- Top Down")
	fmt.Println("Original array:", arr)

	MergeSortWithTopDownInPlace(arr)
	fmt.Println("Sorted array:", arr)
	fmt.Println()
}

func bottomUpMergeSortExample() {
	// Example usage
	arr := []int{9, 3, 7, 5, 1, 8, 2, 6, 10, 4}

	fmt.Println("Merge Sort -- Bottom Up")
	fmt.Println("Original array:", arr)

	MergeSortWithBottomUpInPlace(arr)
	fmt.Println("Sorted array:", arr)
	fmt.Println()
}

func notInplaceMergeSortExample() {
	// Example usage
	arr := []int{9, 3, 7, 5, 1, 8, 2, 6, 10, 4}

	fmt.Println("Merge Sort -- Not in-place Top Down")
	fmt.Println("Original array:", arr)

	result := MergeSortNotInplace(arr)
	fmt.Println("Original array after sort:", arr)
	fmt.Println("Sorted array:", result)
	fmt.Println()
}
