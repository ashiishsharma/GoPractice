package dataStructures

import (
	"errors"
	"iter"
)

type IntegerLinkedList interface {
	int | int64 | float64
}

type ICustomLinkedList[T IntegerLinkedList] interface {
	add(T)
	addToIndex(T, index int) error
	remove(index int) (T, error)
	getSize() int
}
type CustomLinkedList[T IntegerLinkedList] struct {
	size int
	head *Node[T]
}

func NewCustomLinkedList[T IntegerLinkedList]() CustomLinkedList[T] {
	return CustomLinkedList[T]{}
}

func (customLinkedList *CustomLinkedList[T]) add(t T) {
	node := NewNode(t, nil)
	if customLinkedList.head == nil {
		customLinkedList.head = node
	} else {
		iterationNode := customLinkedList.head
		for iterationNode.next != nil {
			iterationNode = iterationNode.next
		}
		iterationNode.next = node
	}
	customLinkedList.size++
}

func (customLinkedList *CustomLinkedList[T]) addToIndex(t T, index int) error {
	var nodeToAdd = NewNode(t, nil)
	var err error
	if index < 0 || index >= customLinkedList.size {
		return errors.New("index out of range")
	} else if customLinkedList.size == 0 {
		return errors.New("Empty List")
	} else if index == 0 {
		nodeToAdd.next = customLinkedList.head
		customLinkedList.head = nodeToAdd
		err = nil
	} else {
		iterationNode := customLinkedList.head
		for i := 0; i < index && i < customLinkedList.size; i++ {
			if i == index-1 {
				break
			}
			iterationNode = iterationNode.next
		}
		if iterationNode.next != nil {
			nodeToAdd.next = iterationNode.next
			iterationNode.next = nodeToAdd
		} else {
			iterationNode.next = nodeToAdd
		}
	}
	if err == nil {
		customLinkedList.size++
	}
	return err
}

func (customLinkedList *CustomLinkedList[T]) remove(index int) (T, error) {
	var elementToRemove T
	var err error
	if index < 0 || index >= customLinkedList.size {
		return elementToRemove, errors.New("index out of range")
	} else if customLinkedList.size == 0 {
		return elementToRemove, errors.New("Empty List")
	} else if customLinkedList.size == 1 && index == 0 {
		elementToRemove = customLinkedList.head.data
		customLinkedList.head = nil
		err = nil
	} else if index == 0 {
		elementToRemove = customLinkedList.head.data
		customLinkedList.head = customLinkedList.head.next
		err = nil
	} else {
		iterationNode := customLinkedList.head
		for i := 0; i < index && i < customLinkedList.size; i++ {
			if i == index-1 {
				break
			}
			iterationNode = iterationNode.next
		}
		elementToRemove = iterationNode.next.data
		if iterationNode.next.next != nil {
			iterationNode.next = iterationNode.next.next
		} else {
			iterationNode.next = nil
		}
	}
	customLinkedList.size--
	return elementToRemove, err
}

func (customLinkedList *CustomLinkedList[T]) getSize() int {
	return customLinkedList.size
}

type Node[T IntegerLinkedList] struct {
	data T
	next *Node[T]
}

func NewNode[T IntegerLinkedList](data T, next *Node[T]) *Node[T] {
	return &Node[T]{data: data, next: next}
}

func (customLinkedList *CustomLinkedList[T]) All() (iteratorSequence iter.Seq[T]) {
	iteratorSequence = func(yield func(T) bool) {
		for node := customLinkedList.head; node != nil; node = node.next {
			if !yield(node.data) {
				return
			}
		}
	}
	return iteratorSequence
}
