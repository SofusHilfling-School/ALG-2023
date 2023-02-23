package main

import (
	"fmt"
)

type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	data int
	next *Node
}

func (list *LinkedList) Add(item int) {
	new := Node{data: item}
	if list.head == nil {
		list.head = &new
		list.tail = &new
	} else {
		list.tail.next = &new
		list.tail = &new
	}
}

func (list *LinkedList) Print() {
	s := "[ "
	n := list.head
	for n != nil {
		s += fmt.Sprintf("%d ", n.data)
		n = n.next
	}
	s += "]"
	fmt.Println(s)
}

func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList) Clear() {
	list.head = nil
	list.tail = nil
}

func (list *LinkedList) Contains(item int) bool {
	n := list.head
	for n != nil {
		if n.data == item {
			return true
		}
		n = n.next
	}
	return false
}

func (list *LinkedList) GetFirst() *Node {
	return list.head
}

func (list *LinkedList) GetLast() *Node {
	return list.tail
}

func (list *LinkedList) RemoveFirst() {
	if list.head != nil {
		list.head = list.head.next
	}
}

func (list *LinkedList) RemoveLast() {
	n := list.head
	if n.next == nil {
		list.head = nil
		list.tail = nil
		return
	}

	for n != nil {
		if n.next.next == nil {
			n.next = nil
			list.tail = n
			n = nil
		} else {
			n = n.next
		}
	}
}

func (list *LinkedList) AddFirst(item int) {
	list.head = &Node{data: item, next: list.head}
}

func (list *LinkedList) AddLast(data int) {
	// n := list.head
	// if n == nil {
	// 	newNode := &Node{data: item}
	// 	list.head = newNode
	// 	list.tail = newNode
	// }

	// for n != nil {
	// 	if n.next == nil {
	// 		newNode := &Node{data: item}
	// 		list.tail = newNode
	// 		n.next = newNode
	// 		n = nil
	// 	} else {
	// 		n = n.next
	// 	}
	// }
	newNode := &Node{data, nil}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}
}

func (list *LinkedList) Reverse() {
	var prev *Node = nil
	n := list.head
	if n != nil {
		next := n.next
		n.next = prev
		prev = n
		n = next
	}
	list.tail = list.head
	list.tail = prev
}

func NewLinkedList() LinkedList {
	return LinkedList{}
}
