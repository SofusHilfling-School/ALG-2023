package main

import "fmt"

func main() {
	BinarySearchTreeExample()
}

func LinearSearchExample() {
	arr := []int{4, 2, 8, 5, 1, 9, 3}
	target := 5

	index := linearSearch(arr, target)
	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}

func BinarySearchExample() {
	// array must be sorted
	arr := []int{1, 2, 3, 4, 5, 8, 9, 14, 19, 24, 25, 50, 88}
	target := 3

	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}

func BinarySearchTreeExample() {
	bst := &BinarySearchTree{}

	// 	      8
	// 	   /    \
	//    3      10
	//  /  \    /  \
	// 1    6  9    14
	//     / \     /
	//    5   7   13

	// Insert values into the BST
	bst.Insert(8)
	bst.Insert(3)
	bst.Insert(10)
	bst.Insert(1)
	bst.Insert(6)
	bst.Insert(5)

	// Search for values in the BST
	fmt.Println(bst.Search(9))  // false
	fmt.Println(bst.Search(2))  // false
	fmt.Println(bst.Search(10)) // true

	bst.Insert(14)
	bst.Insert(7)
	bst.Insert(13)
	bst.Insert(9)

	// Search for values in the BST
	fmt.Println(bst.Search(9)) // true
	fmt.Println(bst.Search(4)) // false
}
