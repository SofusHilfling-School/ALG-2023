package main

import (
	"fmt"
)

type IntArrayList struct {
	data []int
	size int
}

func (list *IntArrayList) Get(index int) int {
	return list.data[index]
}

func (list *IntArrayList) Add(item int) {
	if list.size == len(list.data) {
		list.resize(len(list.data) * 2)
	}
	list.data[list.size] = item
	list.size += 1
}

func (list *IntArrayList) Delete(item int) bool {
	itemIndex := list.find(item)
	if itemIndex < 0 {
		return false
	}
	for i := itemIndex; i+1 < list.size; i++ {
		list.data[i] = list.data[i+1]
	}
	list.data[list.size] = 0
	list.size -= 1
	if list.size > 0 && list.size == len(list.data)/4 {
		list.resize(len(list.data) / 2)
	}
	return true
}

func (list *IntArrayList) Search(item int) bool {
	return list.find(item) > -1
}

func (list *IntArrayList) Set(index, item int) {
	if index > list.size || index < 0 {
		panic("index out of bounds")
	}
	list.data[index] = item
}

func (list *IntArrayList) Insert(index, item int) {
	if list.size == len(list.data) {
		list.resize(len(list.data) * 2)
	}
	for i := list.size + 1; i > index; i-- {
		list.data[i] = list.data[i-1]
	}
	list.data[index] = item
	list.size += 1
}

func (list *IntArrayList) Clear() {
	list.data = make([]int, list.size)
}

func (list *IntArrayList) Size() int {
	return list.size
}

func (list *IntArrayList) IsEmpty() bool {
	return list.size > 0
}

func (list *IntArrayList) ToArray() []int {
	array := make([]int, list.size)
	for i, v := range list.data {
		if i == list.size {
			break
		}
		array[i] = v
	}
	return array
}

func (list *IntArrayList) Reverse() {
	new := make([]int, len(list.data))

	j := 0
	for i := list.size - 1; i >= 0; i-- {
		new[j] = list.data[i]
		j += 1
	}
	list.data = new
}

func (list *IntArrayList) resize(newSize int) {
	fmt.Println("Resizing")
	newData := make([]int, newSize)
	for i, v := range list.data {
		newData[i] = v
	}
	list.data = newData
}

func (list *IntArrayList) find(item int) int {
	for i := 0; i < list.size; i++ {
		if list.data[i] == item {
			return i
		}
	}
	return -1
}

func NewArrayList() IntArrayList {
	return IntArrayList{data: make([]int, 2), size: 0}
}
