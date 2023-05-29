package main

import "fmt"

func main() {
	
}

func constant() {
	arr := []string{"here", "is", "a", "list", "of", "strings"}

	fmt.Println(arr[2])
}

func logarithmic() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	mid := len(arr) / 2
	for mid <= 0 {
		mid = mid / 2
		fmt.Println("", mid)
	}
}

// n
func linear() {
	n := 10
	for i := 0; i < n; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}

// n*2
func linear2() {
	n := 10
	for i := 0; i < n; i++ {
		fmt.Print(i)
	}
	for i := 0; i < n; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}

// n*n or n^2
func quadratic() {
	n := 10
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(i)
		}
	}
	fmt.Println()
}

// n*n or n^2
func quadratic2() {
	n := 10
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(i*j)
		}
	}
	fmt.Println()
	for j := 0; j < n; j++ {
		fmt.Print(j)
	}
	fmt.Println()
}

func quadraticMemory() {
	n := 10
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		tmp := make([]int, n)
		for j := 0; j < n; j++ {
			tmp = append(tmp, j*i)
		}
		result = append(result, tmp)
	}
	fmt.Println(result)
}

// 2^n
func exponential() {
	var fibonacci func(n int) int
	fibonacci := func(n int) int {
		if n <= 1 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}

	fmt.Println(fibonacci(10))
}
