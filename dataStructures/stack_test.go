package dataStructures

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := NewStack[int]()
	var value int
	value = 1
	stack.Push(value)
	var actual int = stack.Pop()

	if value != actual {
		t.Errorf("expected '%d' but got '%d'", actual, value)
	}
}
