package main

import "fmt"

// Stack class
type Stack[T any] []T

// Push adds a value to the stack
func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

// Pop removes and returns a value from the stack in last to first order
func (s *Stack[T]) Pop() (val T) {
	val = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return
}

// Print iterates from the first value to the last value of Stack and prints it
func (s *Stack[T]) Print() {
	for _, val := range *s {
		fmt.Printf("%v - ", val)
	}

	fmt.Println()
}

func main() {
	s := make(Stack[any], 0)

	s.Push(1)
	s.Push("Two")
	s.Push(3.14)
	s.Push(true)

	s.Print()                           // 1 - Two - 3.14 - true
	fmt.Printf("Popped: %v\n", s.Pop()) // Popped: true
	fmt.Printf("Popped: %v\n", s.Pop()) // Popped: 3.14
	s.Print()                           // 1 - Two -
}
