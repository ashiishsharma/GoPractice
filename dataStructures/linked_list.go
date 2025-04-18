package dataStructures

type IntegerLinkedList interface {
	int | int64 | float64
}

type ICustomLinkedList[T IntegerLinkedList] interface {
	add(T)
	remove(int)
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
