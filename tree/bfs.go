package tree

import "reflect"

type queue[T any] []T

func (q *queue[T]) enqueue(data T) {
	*q = append(*q, data)
}

func (q *queue[T]) deque() T {
	data := (*q)[0]
	*q = (*q)[1:]
	return data
}

func BreadthFirstSearch[T any](root *BinaryNode[T], data T) bool {
	q := make(queue[*BinaryNode[T]], 0)
	q.enqueue(root)

	for len(q) > 0 {
		node := q.deque()
		if node == nil {
			continue
		}

		if reflect.DeepEqual(node.Data, data) {
			return true
		}

		q.enqueue(node.Left)
		q.enqueue(node.Right)
	}

	return false
}
