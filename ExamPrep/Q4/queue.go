package main

type Queue struct {
	arr []int
}

func NewQueue(val ...int) *Queue {
	arr := make([]int, len(val))
	copy(arr, val)

	return &Queue{
		arr: arr,
	}
}

func (q *Queue) Dequeue() int {
	itm := q.arr[0]
	q.arr = q.arr[1:]
	return itm
}

func (q *Queue) Enqueue(val int) {
	q.arr = append(q.arr, val)
}

func (q *Queue) Empty() bool {
	return len(q.arr) <= 0
}
