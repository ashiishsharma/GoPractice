package dataStructures

type IntegerStack interface {
	int | float64
}

type IStack[T IntegerStack] interface {
	Push(T)
	Pop() T
	Peek() T
	Size() int
}
type Stack[T IntegerStack] struct {
	data []T
}

func (stack *Stack[T]) Push(i T) {
	stack.data = append(stack.data, i)
}

func (stack *Stack[T]) Pop() T {
	return stack.data[len(stack.data)-1]
}

func (stack *Stack[T]) Peek() T {
	return stack.data[len(stack.data)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func NewStack[T IntegerStack]() *Stack[T] {
	return &Stack[T]{}
}
