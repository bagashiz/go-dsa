package tree

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func BSTDepthFirstSearch[T Number](root *BinaryNode[T], data T) bool {
	return searchBST(root, data)
}

func searchBST[T Number](node *BinaryNode[T], data T) bool {
	if node == nil {
		return false
	}

	if node.Data == data {
		return true
	}

	if node.Data < data {
		return searchBST(node.Right, data)
	}

	return searchBST(node.Left, data)
}
