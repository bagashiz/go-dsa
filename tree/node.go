package tree

type BinaryNode[T any] struct {
	Data  T
	Right *BinaryNode[T]
	Left  *BinaryNode[T]
}
