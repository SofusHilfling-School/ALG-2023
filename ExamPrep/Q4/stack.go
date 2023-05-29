package main

type Stack struct {
	arr []int
}

func NewStack() *Stack {
	return &Stack{arr: make([]int, 0)}
}

func (s *Stack) Push(v int) {
	s.arr = append(s.arr, v)
}

func (s *Stack) Pop() int {

	l := len(s.arr)
	if l <= 0 {
		return 0
	}
	result := s.arr[l-1]
	if l == 1 {
		s.arr = make([]int, 0)
	} else {
		s.arr = s.arr[:l-1]
	}

	return result
}

func (s *Stack) Empty() bool {
	return len(s.arr) <= 0
}
