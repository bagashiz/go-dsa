package sort_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/go-dsa/sort"
)

func TestStackString(t *testing.T) {
	testCases := []struct {
		push   []string
		pop    []string
		peek   string
		length int
	}{
		{push: []string{"apple", "banana", "cherry"}, pop: []string{"cherry", "banana"}, peek: "apple", length: 1},
		{push: []string{"orange"}, pop: []string{}, peek: "orange", length: 1},
		{push: []string{"grape"}, pop: []string{"grape"}, peek: "", length: 0},
		{push: []string{"one", "two", "three", "four", "five"}, pop: []string{"five", "four", "three", "two", "one"}, peek: "", length: 0},
		{push: []string{"zero"}, pop: []string{"zero"}, peek: "", length: 0},
		{push: []string{}, pop: []string{""}, peek: "", length: 0},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			stack := sort.NewStack[string]()

			for _, val := range tc.push {
				stack.Push(val)
			}

			for _, val := range tc.pop {
				popped := stack.Pop()
				if popped != val {
					t.Errorf("[case: %s] want %v, got %v", index, val, popped)
				}
			}

			peeked := stack.Peek()
			if peeked != tc.peek {
				t.Errorf("[case: %s] want %v, got %v", index, tc.peek, peeked)
			}

			length := stack.Length()
			if length != tc.length {
				t.Errorf("[case: %s] want %v, got %v", index, tc.length, length)
			}
		})
	}
}

func TestStackInt(t *testing.T) {
	testCases := []struct {
		push   []int
		pop    []int
		peek   int
		length int
	}{
		{push: []int{1, 2, 3}, pop: []int{3, 2}, peek: 1, length: 1},
		{push: []int{4}, pop: []int{}, peek: 4, length: 1},
		{push: []int{5}, pop: []int{5}, peek: 0, length: 0},
		{push: []int{6, 7, 8, 9, 10}, pop: []int{10, 9, 8, 7, 6}, peek: 0, length: 0},
		{push: []int{0}, pop: []int{0}, peek: 0, length: 0},
		{push: []int{}, pop: []int{0}, peek: 0, length: 0},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			stack := sort.NewStack[int]()

			for _, val := range tc.push {
				stack.Push(val)
			}

			for _, val := range tc.pop {
				popped := stack.Pop()
				if popped != val {
					t.Errorf("[case: %s] want %v, got %v", index, val, popped)
				}
			}

			peeked := stack.Peek()
			if peeked != tc.peek {
				t.Errorf("[case: %s] want %v, got %v", index, tc.peek, peeked)
			}

			length := stack.Length()
			if length != tc.length {
				t.Errorf("[case: %s] want %v, got %v", index, tc.length, length)
			}
		})
	}
}

func TestStackStruct(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	testCases := []struct {
		push   []person
		pop    []person
		peek   person
		length int
	}{
		{push: []person{{name: "Alice", age: 25}, {name: "Bob", age: 30}}, pop: []person{{name: "Bob", age: 30}}, peek: person{name: "Alice", age: 25}, length: 1},
		{push: []person{{name: "Charlie", age: 35}}, pop: []person{}, peek: person{name: "Charlie", age: 35}, length: 1},
		{push: []person{{name: "Dave", age: 40}}, pop: []person{{name: "Dave", age: 40}}, peek: person{}, length: 0},
		{push: []person{{name: "Eve", age: 45}, {name: "Frank", age: 50}, {name: "Grace", age: 55}}, pop: []person{{name: "Grace", age: 55}, {name: "Frank", age: 50}, {name: "Eve", age: 45}}, peek: person{}, length: 0},
		{push: []person{{name: "Harry", age: 60}}, pop: []person{{name: "Harry", age: 60}}, peek: person{}, length: 0},
		{push: []person{}, pop: []person{{}}, peek: person{}, length: 0},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			stack := sort.NewStack[person]()

			for _, val := range tc.push {
				stack.Push(val)
			}

			for _, val := range tc.pop {
				popped := stack.Pop()
				if popped != val {
					t.Errorf("[case: %s] want %v, got %v", index, val, popped)
				}
			}

			peeked := stack.Peek()
			if peeked != tc.peek {
				t.Errorf("[case: %s] want %v, got %v", index, tc.peek, peeked)
			}

			length := stack.Length()
			if length != tc.length {
				t.Errorf("[case: %s] want %v, got %v", index, tc.length, length)
			}
		})
	}
}
