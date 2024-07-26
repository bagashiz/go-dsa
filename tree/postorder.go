package tree

func PostOrderTraverse[T any](root *BinaryNode[T]) []T {
	path := make([]T, 0)
	postOrderTraverse(root, &path)
	return path
}

func postOrderTraverse[T any](node *BinaryNode[T], path *[]T) {
	if node == nil {
		return
	}

	postOrderTraverse(node.Left, path)
	postOrderTraverse(node.Right, path)
	*path = append(*path, node.Data)
}
