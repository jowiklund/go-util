package queue

import "fmt"

type Queue[T any] struct {
	items []T
}

func NewQueue[T any](size int) Queue[T] {
	slice := make([]T, size)
	slice = slice[:0]
	return Queue[T]{
		items: slice,
	}
}

func (q *Queue[T]) Enqueue(item T) error {
	if q.Full() {
		return fmt.Errorf("Queue is full")
	}
	q.items = append(q.items, item)
	return nil
}

func (q *Queue[T]) Dequeue() (T, error) {
	var res T
	if q.Empty() {
		return res, fmt.Errorf("Queue is empty")
	}
	res = q.items[0]
	q.items = q.items[1:]
	return res, nil
}

func (q *Queue[T]) Peek() (T, bool) {
	var item T
	if q.Empty() {
		return item, false
	}
	return q.items[0], true
}

func (q *Queue[T]) Empty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Full() bool {
	return len(q.items) == cap(q.items)
}

