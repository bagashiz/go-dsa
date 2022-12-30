package main

import "fmt"

// Node class
type Node struct {
	val  any
	next *Node
	prev *Node
}

// DoublyLinkedList class
type DoublyLinkedList struct {
	head *Node
}

// AddFirst adds the node to the start of the linked list
func (dll *DoublyLinkedList) AddFirst(val any) {
	node := &Node{}
	node.val = val
	node.next = nil

	if dll.head != nil {
		node.next = dll.head
		dll.head.prev = node
	}

	dll.head = node
}

// AddLast adds the node at the end of the linked list
func (dll *DoublyLinkedList) AddLast(val any) {
	node := &Node{}
	node.val = val
	node.next = nil

	tail := dll.Tail()

	if tail != nil {
		tail.next = node
		node.prev = tail
	}
}

// AddAfter adds the node after a specific node
func (dll *DoublyLinkedList) AddAfter(prevNodeVal any, val any) {
	node := &Node{}
	node.val = val
	node.next = nil

	var nodeWith = &Node{}
	nodeWith = dll.NodeWithValue(prevNodeVal)

	if nodeWith != nil {
		node.next = nodeWith.next
		node.prev = nodeWith
		nodeWith.next = node
	}
}

// Tail returns the node at the end of the linked list
func (dll *DoublyLinkedList) Tail() (tail *Node) {
	for node := dll.head; node != nil; node = node.next {
		if node.next == nil {
			tail = node
		}
	}

	return
}

// NodeWithValue returns the node with the given value
func (dll *DoublyLinkedList) NodeWithValue(val any) (nodeWith *Node) {
	for node := dll.head; node != nil; node = node.next {
		if node.val == val {
			nodeWith = node
			return
		}
	}

	return nil
}

// NodeBetweenValues returns the node that has value between the previous and next node's given values
func (dll *DoublyLinkedList) NodeBetweenValues(prevVal, nextVal any) (nodeWith *Node) {
	for node := dll.head; node != nil; node = node.next {
		if node.prev != nil && node.next != nil {
			if node.prev.val == prevVal && node.next.val == nextVal {
				nodeWith = node
				return
			}
		}
	}

	return nil
}

// Print iterates from the head node to the last node and prints the values
func (dll *DoublyLinkedList) Print() {
	for node := dll.head; node != nil; node = node.next {
		fmt.Printf("%v <-> ", node.val)
	}

	fmt.Println()
}

func main() {
	dll := DoublyLinkedList{}
	dll.AddFirst(1)       // 1 <->
	dll.AddFirst("Three") // Three <-> 1 <->
	dll.AddLast(false)    // Three <-> 1 <-> false <->
	dll.AddAfter(1, 3.14) // Three <-> 1 <-> 3.14 <-> false <->
	dll.Print()

	node := &Node{}
	node = dll.NodeBetweenValues(1, false) // 3,14
	fmt.Println(node.val)
}
