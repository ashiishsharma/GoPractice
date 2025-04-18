package dataStructures

import "errors"

type IntegerQueue interface {
	int | float64
}

type IQueue[T IntegerQueue] interface {
	Enqueue(T)
	Dequeue() (T, error)
	size() int
}
type Queue[T IntegerQueue] struct {
	data []T
}

func (queue *Queue[T]) Enqueue(t T) {
	queue.data = append(queue.data, t)
}

func (queue *Queue[T]) Dequeue() (T, error) {
	if len(queue.data) == 0 {
		var zeroValue T
		return zeroValue, errors.New("empty queue")
	} else if len(queue.data) == 1 {
		var val = queue.data[0]
		var emptyArray []T
		queue.data = emptyArray
		return val, nil
	}
	var val = queue.data[0]
	queue.data = queue.data[1:]
	return val, nil
}

func (queue *Queue[T]) size() int {
	return len(queue.data)
}

func NewQueue[T IntegerQueue]() *Queue[T] {
	return &Queue[T]{}
}
