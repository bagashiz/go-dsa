package main

import "fmt"

// Node class
type Node struct {
	val  any
	next *Node
}

// SinglyLinkedList class
type SinglyLinkedList struct {
	head *Node
}

// AddFirst adds the node to the start of the linked list
func (sll *SinglyLinkedList) AddFirst(val any) {
	node := Node{}
	node.val = val

	if sll.head != nil {
		node.next = sll.head
	}

	sll.head = &node
}

// AddLast adds the node at the end of the linked list
func (sll *SinglyLinkedList) AddLast(val any) {
	node := &Node{}
	node.val = val
	node.next = nil

	tail := sll.Tail()

	if tail != nil {
		tail.next = node
	}
}

// AddAfter adds the node after a specific node
func (sll *SinglyLinkedList) AddAfter(prevNodeVal any, val any) {
	node := &Node{}
	node.val = val
	node.next = nil

	var nodeWith = &Node{}
	nodeWith = sll.NodeWithValue(prevNodeVal)

	if nodeWith != nil {
		node.next = nodeWith.next
		nodeWith.next = node
	}
}

// Tail returns the node at the end of the linked list
func (sll *SinglyLinkedList) Tail() (tail *Node) {
	for node := sll.head; node != nil; node = node.next {
		if node.next == nil {
			tail = node
		}
	}

	return
}

// NodeWithValue returns the node with the given value
func (sll *SinglyLinkedList) NodeWithValue(val any) (nodeWith *Node) {
	for node := sll.head; node != nil; node = node.next {
		if node.val == val {
			nodeWith = node
			return
		}
	}

	return nil
}

// Print iterates from the head node to the last node and prints the values
func (sll *SinglyLinkedList) Print() {
	for node := sll.head; node != nil; node = node.next {
		fmt.Printf("%v -> ", node.val)
	}

	fmt.Println()
}

func main() {
	sll := SinglyLinkedList{}
	sll.AddFirst(1)       // 1 ->
	sll.AddFirst("Three") // Three -> 1 ->
	sll.AddLast(false)    // Three -> 1 -> false ->
	sll.AddAfter(1, 3.14) // Three -> 1 -> 3.14 -> false ->
	sll.Print()
}
