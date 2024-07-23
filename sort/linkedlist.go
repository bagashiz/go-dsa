package sort

import "reflect"

type SLNode[T any] struct {
	data T
	next *SLNode[T]
}

type SinglyLinkedList[T any] struct {
	head   *SLNode[T]
	tail   *SLNode[T]
	length int
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (sll *SinglyLinkedList[T]) Length() int {
	return sll.length
}

func (sll *SinglyLinkedList[T]) InsertAt(data T, index int) {
	if index < 0 || index > sll.length {
		return
	}

	if index == 0 {
		sll.Prepend(data)
		return
	}

	if index == sll.length {
		sll.Append(data)
		return
	}

	node := &SLNode[T]{data: data}

	current := sll.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	node.next = current.next
	current.next = node

	sll.length++
}

func (sll *SinglyLinkedList[T]) Remove(data T) T {
	if sll.head == nil {
		return *new(T)
	}

	if reflect.DeepEqual(sll.head.data, data) {
		return sll.RemoveAt(0)
	}

	current := sll.head
	for current.next != nil {
		if reflect.DeepEqual(current.next.data, data) {
			current.next = current.next.next
			sll.length--
			return data
		}
		current = current.next
	}

	return *new(T)
}

func (sll *SinglyLinkedList[T]) RemoveAt(index int) T {
	if index < 0 || index >= sll.length {
		return *new(T)
	}

	sll.length--

	if index == 0 {
		data := sll.head.data
		sll.head = sll.head.next
		return data
	}

	current := sll.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	data := current.next.data
	current.next = current.next.next

	return data
}

func (sll *SinglyLinkedList[T]) Append(data T) {
	node := &SLNode[T]{data: data}

	sll.length++

	if sll.head == nil {
		sll.head = node
		sll.tail = node
		return
	}

	tail := sll.tail
	tail.next = node
	sll.tail = node
}

func (sll *SinglyLinkedList[T]) Prepend(data T) {
	node := &SLNode[T]{data: data}

	sll.length++

	if sll.head == nil {
		sll.head = node
		sll.tail = node
		return
	}

	head := sll.head
	node.next = head
	sll.head = node
}

func (sll *SinglyLinkedList[T]) Get(index int) T {
	if index < 0 || index > sll.length {
		return *new(T)
	}

	current := sll.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	if current == nil {
		return *new(T)
	}

	return current.data
}
