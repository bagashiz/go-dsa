package main

import "fmt"

// PriorityQueue is an array of Elements type with priority
type PriorityQueue[T any] []*Element[T]

// Element class
type Element[T any] struct {
	priority int
	val      T
}

// New initializes the property of the Element class
func (el *Element[T]) New(priority int, val T) {
	el.priority = priority
	el.val = val
}

// Enque adds Element object to PriorityQueue based on the priority
func (pq *PriorityQueue[T]) Enque(el *Element[T]) {
	if len(*pq) == 0 {
		*pq = append(*pq, el)
	} else {
		appended := false

		for i, addedElement := range *pq {
			if el.priority < addedElement.priority {
				*pq = append((*pq)[:i], append(PriorityQueue[T]{el}, (*pq)[i:]...)...)
				appended = true
				break
			}
		}

		if !appended {
			*pq = append(*pq, el)
		}
	}
}

// Deque returns the first element of PriorityQueue and removes it from PriorityQueue
func (pq *PriorityQueue[T]) Deque() (first T) {
	first = (*pq)[0].val
	*pq = (*pq)[1:]
	return
}

// Print iterates from the first element to the last element of PriorityQueue and prints its value
func (pq *PriorityQueue[T]) Print() {
	for _, el := range *pq {
		fmt.Printf("%v - ", el.val)
	}

	fmt.Println()
}

func main() {
	pq := make(PriorityQueue[any], 0)

	pq.Enque(&Element[any]{
		priority: 420,
		val:      "Money",
	})

	pq.Enque(&Element[any]{
		priority: 69,
		val:      "Love",
	})

	pq.Enque(&Element[any]{
		priority: 100,
		val:      "Power",
	})

	pq.Print()                             // Love - Power - Money -
	fmt.Printf("Dequed: %v\n", pq.Deque()) // Dequed: Love
	pq.Print()                             // Power - Money -

}
