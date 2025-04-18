package dataStructures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_EnqueueDeque(t *testing.T) {
	queue := NewQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	var dequedValue int
	err := error(nil)
	dequedValue, err = queue.Dequeue()
	assert.Equal(t, 1, dequedValue)

	dequedValue, err = queue.Dequeue()
	assert.Equal(t, 2, dequedValue)

	assert.Nil(t, err)
}

func TestQueueSize(t *testing.T) {
	queue := NewQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	assert.Equal(t, 4, queue.size())
	queue.Dequeue()
	assert.Equal(t, 3, queue.size())
}
