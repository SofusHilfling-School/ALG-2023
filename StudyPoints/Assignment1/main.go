package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Bubble sort:")
	arr := []int{9823, 328, 2349, 49, 3, 52, 2}
	bubbleSort(arr)
	fmt.Println(arr)

	fmt.Println("\nStack:")
	stack := Stack{}
	stack.push(5)
	stack.push(2)
	stack.print()
	val, _ := stack.pop()
	fmt.Println(val)
	val, _ = stack.peek()
	fmt.Println(val)
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// could also write it like this
				// tmp := arr[j]
				// arr[j] = arr[j+1]
				// arr[j+1] = tmp
			}
		}
	}
}

type Node struct {
	data int
	next *Node
}

type Stack struct {
	head *Node
	size int
}

func (s *Stack) pop() (int, error) {
	if s.head == nil {
		return 0, errors.New("empty stack")
	}
	dat := s.head.data
	s.head = s.head.next
	s.size--
	return dat, nil
}

func (s *Stack) push(item int) {
	s.head = &Node{data: item, next: s.head}
	s.size++
}

func (s *Stack) peek() (int, error) {
	if s.head == nil {
		return 0, errors.New("empty stack")
	}
	return s.head.data, nil
}

// look for an element in the stack and returns how far the item is from the top
// returns -1 if no item match
func (s *Stack) find(item int) int {
	i := 0
	n := s.head
	for n != nil {
		if n.data == item {
			return i
		}
		i++
		n = n.next
	}
	return -1
}

func (s *Stack) print() {
	n := s.head
	result := "[ "
	for n != nil {
		result += fmt.Sprintf("%d ", n.data)
		n = n.next
	}
	result += "]"
	fmt.Println(result)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func binarySearch(arr []int, target int) bool {
	return binarySearchHelper(arr, target, 0, len(arr)-1)
}

func binarySearchHelper(arr []int, target int, left int, right int) bool {
	if left > right {
		return false
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return true
	} else if arr[mid] > target {
		return binarySearchHelper(arr, target, left, mid-1)
	} else {
		return binarySearchHelper(arr, target, mid+1, right)
	}
}
