package main

import "fmt"

// Queue is an array of Elements type
type Queue[T any] []T

// Enque adds Element object to Queue
func (q *Queue[T]) Enque(val T) {
	*q = append(*q, val)
}

// Deque returns the first valement of Queue and removes it from PriorityQueue
func (q *Queue[T]) Deque() (first T) {
	first = (*q)[0]
	*q = (*q)[1:]
	return
}

// Print iterates from the first element to the last element of Queue and prints its value
func (q *Queue[T]) Print() {
	for _, val := range *q {
		fmt.Printf("%v - ", val)
	}

	fmt.Println()
}

func main() {
	q := make(Queue[any], 0)

	q.Enque("Money")
	q.Enque("Love")
	q.Enque("Power")

	q.Print()                             // Money - Love - Power -
	fmt.Printf("Dequed: %v\n", q.Deque()) // Dequed: Money
	q.Print()                             // Love - Power -

}
