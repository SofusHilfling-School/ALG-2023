package main

import "fmt"

func main() {
	runLinkedList()
}

func runArrayList() {
	list := NewArrayList()
	fmt.Printf("list length: %d, internal size: %d\n", list.size, len(list.data))
	fmt.Println(list.data)
	list.Add(5)
	list.Add(39)
	fmt.Printf("list length: %d, internal size: %d\n", list.size, len(list.data))
	fmt.Println(list.data)
	list.Add(88)
	fmt.Printf("list length: %d, internal size: %d\n", list.size, len(list.data))
	fmt.Println(list.data)
}

func runLinkedList() {
	list := NewLinkedList()
	list.Add(5)
	list.Add(10)
	list.Print()
}
