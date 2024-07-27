package heap

type MinHeap struct {
	data   []int
	length int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) Length() int {
	return h.length
}

func (h *MinHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(h.length)
	h.length++
}

func (h *MinHeap) Delete() int {
	if h.length == 0 {
		return -1
	}

	value := h.data[0]

	h.length--

	if h.length == 0 {
		h.data = []int{}
		return value
	}

	h.data[0] = h.data[h.length]
	h.heapifyDown(0)

	return value
}

func (h *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	parent := h.parent(idx)
	parentValue := h.data[parent]
	currentValue := h.data[idx]

	if parentValue > currentValue {
		h.data[idx], h.data[parent] = parentValue, currentValue
		h.heapifyUp(parent)
	}
}

func (h *MinHeap) heapifyDown(idx int) {
	leftIdx := h.leftChild(idx)
	rightIdx := h.rightChild(idx)

	if idx >= h.length || leftIdx >= h.length {
		return
	}

	leftValue := h.data[leftIdx]
	rightValue := h.data[rightIdx]
	currentValue := h.data[idx]

	if leftValue > rightValue && currentValue > rightValue {
		h.data[idx], h.data[rightIdx] = rightValue, currentValue
		h.heapifyDown(rightIdx)
	} else if rightValue > leftValue && currentValue > leftValue {
		h.data[idx], h.data[leftIdx] = leftValue, currentValue
		h.heapifyDown(leftIdx)
	}
}

func (h *MinHeap) parent(idx int) int {
	return (idx - 1) / 2
}

func (h *MinHeap) leftChild(idx int) int {
	return (idx * 2) + 1
}

func (h *MinHeap) rightChild(idx int) int {
	return (idx * 2) + 2
}
