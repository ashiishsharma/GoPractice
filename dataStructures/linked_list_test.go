package dataStructures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListAdd(t *testing.T) {
	var linkedList = NewCustomLinkedList[int]()
	linkedList.add(34)
	linkedList.add(56)
	linkedList.add(83)
	linkedList.add(97)
	assert.Equal(t, linkedList.getSize(), 4)
}

func TestLinkedListRemove(t *testing.T) {
	var linkedList = NewCustomLinkedList[int]()
	linkedList.add(34)
	linkedList.add(56)
	linkedList.add(83)
	linkedList.add(97)
	assert.Equal(t, linkedList.getSize(), 4)
}
