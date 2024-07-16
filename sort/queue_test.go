package sort_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/go-dsa/sort"
)

func TestQueueString(t *testing.T) {
	testCases := []struct {
		enqueue []string
		deque   []string
		peek    string
		length  int
	}{
		{enqueue: []string{"apple", "banana", "cherry"}, deque: []string{"apple"}, peek: "banana", length: 2},
		{enqueue: []string{"orange"}, deque: []string{}, peek: "orange", length: 1},
		{enqueue: []string{"grape"}, deque: []string{"grape"}, peek: "", length: 0},
		{enqueue: []string{"one", "two", "three", "four", "five"}, deque: []string{"one", "two", "three", "four", "five"}, peek: "", length: 0},
		{enqueue: []string{"zero"}, deque: []string{"zero"}, peek: "", length: 0},
		{enqueue: []string{}, deque: []string{""}, peek: "", length: 0},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			queue := sort.NewQueue[string]()

			for _, val := range tc.enqueue {
				queue.Enqueue(val)
			}

			for _, val := range tc.deque {
				dequeued := queue.Deque()
				if dequeued != val {
					t.Errorf("[case: %s] want %v, got %v", index, val, dequeued)
				}
			}

			peeked := queue.Peek()
			if peeked != tc.peek {
				t.Errorf("[case: %s] want %v, got %v", index, tc.peek, peeked)
			}

			length := queue.Length()
			if length != tc.length {
				t.Errorf("[case: %s] want %v, got %v", index, tc.length, length)
			}
		})
	}
}

func TestQueueInt(t *testing.T) {
	testCases := []struct {
		enqueue []int
		deque   []int
		peek    int
		length  int
	}{
		{enqueue: []int{1, 2, 3}, deque: []int{1}, peek: 2, length: 2},
		{enqueue: []int{4}, deque: []int{}, peek: 4, length: 1},
		{enqueue: []int{5}, deque: []int{5}, peek: 0, length: 0},
		{enqueue: []int{6, 7, 8, 9, 10}, deque: []int{6, 7, 8, 9, 10}, peek: 0, length: 0},
		{enqueue: []int{0}, deque: []int{0}, peek: 0, length: 0},
		{enqueue: []int{}, deque: []int{0}, peek: 0, length: 0},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			queue := sort.NewQueue[int]()

			for _, val := range tc.enqueue {
				queue.Enqueue(val)
			}

			for _, val := range tc.deque {
				dequeued := queue.Deque()
				if dequeued != val {
					t.Errorf("[case: %s] want %v, got %v", index, val, dequeued)
				}
			}

			peeked := queue.Peek()
			if peeked != tc.peek {
				t.Errorf("[case: %s] want %v, got %v", index, tc.peek, peeked)
			}

			length := queue.Length()
			if length != tc.length {
				t.Errorf("[case: %s] want %v, got %v", index, tc.length, length)
			}
		})
	}
}

func TestQueueStruct(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	testCases := []struct {
		enqueue []person
		deque   []person
		peek    person
		length  int
	}{
		{enqueue: []person{{name: "Alice", age: 25}, {name: "Bob", age: 30}}, deque: []person{{name: "Alice", age: 25}}, peek: person{name: "Bob", age: 30}, length: 1},
		{enqueue: []person{{name: "Charlie", age: 35}}, deque: []person{}, peek: person{name: "Charlie", age: 35}, length: 1},
		{enqueue: []person{{name: "Dave", age: 40}}, deque: []person{{name: "Dave", age: 40}}, peek: person{}, length: 0},
		{enqueue: []person{{name: "Eve", age: 45}, {name: "Frank", age: 50}, {name: "Grace", age: 55}}, deque: []person{{name: "Eve", age: 45}, {name: "Frank", age: 50}, {name: "Grace", age: 55}}, peek: person{}, length: 0},
		{enqueue: []person{{name: "Harry", age: 60}}, deque: []person{{name: "Harry", age: 60}}, peek: person{}, length: 0},
		{enqueue: []person{}, deque: []person{{}}, peek: person{}, length: 0},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			queue := sort.NewQueue[person]()

			for _, val := range tc.enqueue {
				queue.Enqueue(val)
			}

			for _, val := range tc.deque {
				dequeued := queue.Deque()
				if dequeued != val {
					t.Errorf("[case: %s] want %v, got %v", index, val, dequeued)
				}
			}

			peeked := queue.Peek()
			if peeked != tc.peek {
				t.Errorf("[case: %s] want %v, got %v", index, tc.peek, peeked)
			}

			length := queue.Length()
			if length != tc.length {
				t.Errorf("[case: %s] want %v, got %v", index, tc.length, length)
			}
		})
	}
}
