package dataStructures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackPushAndPop(t *testing.T) {
	stack := NewStack[int]()
	var value int
	value = 1
	stack.Push(value)
	var actual int = stack.Pop()

	assert.Equal(t, value, actual)
}

func TestStackPeek(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	actual := stack.Peek()
	assert.Equal(t, 3, actual)
}

func TestStackSize(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	actual := stack.Size()
	assert.Equal(t, 3, actual)
}
