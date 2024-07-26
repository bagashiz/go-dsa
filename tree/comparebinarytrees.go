package tree

import "reflect"

func CompareBT[T any](a, b *BinaryNode[T]) bool {
	// structural check
	if a == nil && b == nil {
		return true
	}

	// structural check
	if a == nil || b == nil {
		return false
	}

	// value check
	if !reflect.DeepEqual(a.Data, b.Data) {
		return false
	}

	return CompareBT(a.Left, b.Left) && CompareBT(a.Right, b.Right)
}
