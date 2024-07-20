package sort

type SNode[T any] struct {
	data T
	prev *SNode[T]
}

type Stack[T any] struct {
	head   *SNode[T]
	length int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Length() int {
	return s.length
}

func (s *Stack[T]) Push(data T) {
	node := &SNode[T]{data: data}

	s.length++

	if s.head == nil {
		s.head = node
		return
	}

	node.prev = s.head
	s.head = node
}

func (s *Stack[T]) Pop() T {
	if s.head == nil {
		return *new(T)
	}

	s.length--

	current := s.head
	s.head = s.head.prev

	return current.data
}

func (s *Stack[T]) Peek() T {
	if s.head == nil {
		return *new(T)
	}
	return s.head.data
}
