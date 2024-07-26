package tree

func PreOrderTraverse[T any](root *BinaryNode[T]) []T {
	path := make([]T, 0)
	preOrderTraverse(root, &path)
	return path
}

func preOrderTraverse[T any](node *BinaryNode[T], path *[]T) {
	if node == nil {
		return
	}

	*path = append(*path, node.Data)
	preOrderTraverse(node.Left, path)
	preOrderTraverse(node.Right, path)
}
