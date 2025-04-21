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

func TestLinkedListAddAtIndex(t *testing.T) {
	var linkedList = NewCustomLinkedList[int]()
	linkedList.add(34)
	linkedList.add(56)
	linkedList.add(83)
	linkedList.add(97)
	assert.Equal(t, linkedList.getSize(), 4)

	linkedList.addToIndex(89, 0)
	assert.Equal(t, linkedList.getSize(), 5)

	linkedList.addToIndex(172, 4)
	assert.Equal(t, linkedList.getSize(), 6)

	linkedList.addToIndex(17, 3)
	assert.Equal(t, linkedList.getSize(), 7)

	for v := range linkedList.All() {
		println(v)
	}
}

func TestLinkedListRemove(t *testing.T) {
	var linkedList = NewCustomLinkedList[int]()
	linkedList.add(34)
	linkedList.add(56)
	linkedList.add(83)
	linkedList.add(97)
	linkedList.add(98)
	linkedList.add(992)
	assert.Equal(t, linkedList.getSize(), 6)

	removedElement, err := linkedList.remove(0)
	assert.Nil(t, err)
	assert.Equal(t, 34, removedElement)

	removedElement, err = linkedList.remove(3)
	assert.Nil(t, err)
	assert.Equal(t, 98, removedElement)

	linkedList.add(9)
	linkedList.add(8)
	linkedList.add(92)
	assert.Equal(t, linkedList.getSize(), 7)

	removedElement, err = linkedList.remove(6)
	assert.Nil(t, err)
	assert.Equal(t, 92, removedElement)

	for v := range linkedList.All() {
		println(v)
	}
}
