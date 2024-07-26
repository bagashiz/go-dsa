package tree

func InOrderTraverse[T any](root *BinaryNode[T]) []T {
	path := make([]T, 0)
	inOrderTraverse(root, &path)
	return path
}

func inOrderTraverse[T any](node *BinaryNode[T], path *[]T) {
	if node == nil {
		return
	}

	inOrderTraverse(node.Left, path)
	*path = append(*path, node.Data)
	inOrderTraverse(node.Right, path)
}
