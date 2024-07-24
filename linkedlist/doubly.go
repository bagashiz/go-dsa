package linkedlist

import "reflect"

type DLNode[T any] struct {
	data T
	next *DLNode[T]
	prev *DLNode[T]
}

type DoublyLinkedList[T any] struct {
	head   *DLNode[T]
	tail   *DLNode[T]
	length int
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (dll *DoublyLinkedList[T]) Length() int {
	return dll.length
}

func (dll *DoublyLinkedList[T]) InsertAt(data T, index int) {
	if index < 0 || index > dll.length {
		return
	}

	if index == 0 {
		dll.Prepend(data)
		return
	}

	if index == dll.length {
		dll.Append(data)
		return
	}

	dll.length++

	current := dll.getAt(index)
	node := &DLNode[T]{data: data}

	node.next = current
	node.prev = current.prev
	current.prev = node

	if node.prev != nil {
		node.prev.next = node
	}
}

func (dll *DoublyLinkedList[T]) Append(data T) {
	dll.length++

	node := &DLNode[T]{data: data}

	if dll.tail == nil {
		dll.head = node
		dll.tail = node
		return
	}

	node.prev = dll.tail
	dll.tail.next = node
	dll.tail = node
}

func (dll *DoublyLinkedList[T]) Prepend(data T) {
	dll.length++

	node := &DLNode[T]{data: data}

	if dll.head == nil {
		dll.head = node
		dll.tail = node
		return
	}

	node.next = dll.head
	dll.head.prev = node
	dll.head = node
}

func (dll *DoublyLinkedList[T]) Remove(data T) T {
	if dll.head == nil {
		return *new(T)
	}

	if reflect.DeepEqual(dll.head.data, data) {
		return dll.RemoveAt(0)
	}

	current := dll.head
	for current != nil {
		if reflect.DeepEqual(current.data, data) {
			return dll.removeNode(current)
		}
		current = current.next
	}

	return *new(T)
}

func (dll *DoublyLinkedList[T]) RemoveAt(index int) T {
	if index < 0 || index >= dll.length {
		return *new(T)
	}

	current := dll.getAt(index)

	return dll.removeNode(current)
}

func (dll *DoublyLinkedList[T]) Get(index int) T {
	if index < 0 || index > dll.length {
		return *new(T)
	}

	current := dll.getAt(index)
	if current == nil {
		return *new(T)
	}

	return current.data
}

func (dll *DoublyLinkedList[T]) getAt(index int) *DLNode[T] {
	current := dll.head
	for range index {
		current = current.next
	}
	return current
}

func (dll *DoublyLinkedList[T]) removeNode(node *DLNode[T]) T {
	dll.length--

	if dll.length == 0 {
		dll.head = nil
		dll.tail = nil
		return *new(T)
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node == dll.head {
		dll.head = node.next
	}
	if node == dll.tail {
		dll.tail = node.prev
	}
	return node.data
}
