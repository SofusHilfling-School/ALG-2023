package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	arr := []string{"something", "25", "4", "i", "hej", "all"}
	extraAsignment(arr)
	fmt.Println(arr)
}

func extraAsignment(arr []string) {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if compareString(arr[j-1], arr[j]) {
				tmp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = tmp
			}
		}
	}
}

func compareString(s1 string, s2 string) bool {
	if num1, err1 := strconv.Atoi(s1); err1 == nil {
		if num2, err2 := strconv.Atoi(s2); err2 == nil {
			return num1 > num2
		} else {
			return false
		}
	}

	if len(s1) > len(s2) {
		return true
	} else if len(s1) < len(s2) {
		return false
	}

	c1 := int(s1[0])
	c2 := int(s2[0])

	return c1 > c2
}

func bubbleSort() {

	rand.Seed(time.Now().Unix())
	//Generate a random array of length n
	randArr := rand.Perm(100000)
	start := time.Now()

	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	arr = randArr

	for i := 0; i < len(arr); i++ {
		flag := true
		for j := 1; j < len(arr); j++ {
			if arr[j-1] > arr[j] {
				tmp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = tmp
				flag = false
			}
		}
		//fmt.Println(arr)
		if flag {
			break
		}

		if i%1000 == 0 {
			fmt.Printf("%d, time elapsed: %s\n", i, time.Since(start))
			start = time.Now()
		}
	}
	println("Done!")
}
