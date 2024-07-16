package sort

type QNode[T any] struct {
	data T
	next *QNode[T]
}

type Queue[T any] struct {
	length int
	head   *QNode[T]
	tail   *QNode[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Length() int {
	return q.length
}

func (q *Queue[T]) Enqueue(data T) {
	node := &QNode[T]{data: data}

	q.length++

	if q.tail == nil {
		q.head = node
		q.tail = node
		return
	}

	q.tail.next = node
	q.tail = node
}

func (q *Queue[T]) Deque() T {
	if q.head == nil {
		return *new(T)
	}

	q.length--

	current := q.head
	q.head = q.head.next

	return current.data
}

func (q *Queue[T]) Peek() T {
	if q.head == nil {
		return *new(T)
	}
	return q.head.data
}
