package main

import (
	"fmt"
	"sync"
)

// TreeNode class
type Node[T any] struct {
	key   int
	val   T
	left  *Node[T]
	right *Node[T]
}

// BinarySearchTree class
type BinarySearchTree[T any] struct {
	root *Node[T]
	lock sync.RWMutex
}

// Insert inserts the node with the given key and value to the binary search tree
func (bst *BinarySearchTree[T]) Insert(key int, val T) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	node := &Node[T]{key, val, nil, nil}

	if bst.root == nil {
		bst.root = node
	} else {
		insertTreeNode(bst.root, node)
	}
}

// insertTreeNode inserts the new tree node to the binary search tree
func insertTreeNode[T any](root, node *Node[T]) {
	if node.key < root.key {
		if root.left == nil {
			root.left = node
		} else {
			insertTreeNode(root.left, node)
		}
	} else {
		if root.right == nil {
			root.right = node
		} else {
			insertTreeNode(root.right, node)
		}
	}
}

// InOrder visits all nodes with in order traversing
func (bst *BinarySearchTree[T]) InOrder(f func(T)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	inOrderTree(bst.root, f)
}

// inOrderTree traverses the left, the root, and the right tree
func inOrderTree[T any](node *Node[T], f func(T)) {
	if node != nil {
		inOrderTree(node.left, f)
		f(node.val)
		inOrderTree(node.right, f)
	}
}

// PreOrder visits all nodes with pre-order traversing
func (bst *BinarySearchTree[T]) PreOrder(f func(T)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	preOrderTree(bst.root, f)
}

// preOrderTree traverses the root, the left, and the right tree
func preOrderTree[T any](node *Node[T], f func(T)) {
	if node != nil {
		f(node.val)
		inOrderTree(node.left, f)
		inOrderTree(node.right, f)
	}
}

// PostOrder visits all nodes with post-order traversing
func (bst *BinarySearchTree[T]) PostOrder(f func(T)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	postOrderTree(bst.root, f)
}

// postOrderTree traverses the left, the right, and the root tree
func postOrderTree[T any](node *Node[T], f func(T)) {
	if node != nil {
		inOrderTree(node.left, f)
		inOrderTree(node.right, f)
		f(node.val)
	}
}

// Min finds the node with the minimum value in the binary search tree
func (bst *BinarySearchTree[T]) Min() *T {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	node := bst.root

	if node == nil {
		return nil
	}

	for {
		if node.left == nil {
			return &node.val
		}

		node = node.left
	}
}

// Max finds the node with the maximum value in the binary search tree
func (bst *BinarySearchTree[T]) Max() *T {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	node := bst.root

	if node == nil {
		return nil
	}

	for {
		if node.right == nil {
			return &node.val
		}

		node = node.right
	}
}

// Search searches the specified node in the binary search tree and returns its value
func (bst *BinarySearchTree[T]) Search(key int) *T {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	return searchTreeNode(bst.root, key)
}

// searchTreeNode searches the specified node in the binary search tree and returns its value
func searchTreeNode[T any](node *Node[T], key int) *T {
	if node == nil {
		return nil
	}

	if key < node.key {
		return searchTreeNode(node.left, key)
	}

	if key > node.key {
		return searchTreeNode(node.right, key)
	}

	return &node.val
}

// Remove removes the node with key that is passed in
func (bst *BinarySearchTree[T]) Remove(key int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	removeTreeNode(bst.root, key)
}

// removeTreeNode removes the node with key that is passed in
func removeTreeNode[T any](node *Node[T], key int) *Node[T] {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = removeTreeNode(node.left, key)
		return node
	}
	if key > node.key {
		node.right = removeTreeNode(node.right, key)
		return node
	}

	// key == node.key
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}

	if node.left == nil {
		node = node.right
		return node
	}

	if node.right == nil {
		node = node.left
		return node
	}

	leftmostRightNode := node.right
	for {
		// find the smallest value on the right side
		if leftmostRightNode != nil && leftmostRightNode.left != nil {
			leftmostRightNode = leftmostRightNode.left
		} else {
			break
		}
	}

	node.key, node.val = leftmostRightNode.key, leftmostRightNode.val
	node.right = removeTreeNode(node.right, node.key)
	return node
}

// String prints the tree into a string format
func (bst *BinarySearchTree[T]) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	stringify(bst.root, 0)
}

// stringify prints the tree into a string format
func stringify[T any](node *Node[T], level int) {
	if node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "		"
		}
		format += "---> "
		level++
		stringify(node.left, level)
		fmt.Printf(format+"(%v)[%v]\n", node.key, node.val)
		stringify(node.right, level)
	}
}

func main() {
	bst := &BinarySearchTree[any]{}

	bst.Insert(8, "Eight")
	bst.Insert(3, "Three")
	bst.Insert(10, "Ten")
	bst.Insert(1, "One")
	bst.Insert(6, "Six")

	bst.String()

	bst.Remove(3)

	bst.String()
}
