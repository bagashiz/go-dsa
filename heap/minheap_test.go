package heap_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/go-dsa/heap"
)

func TestMinHeap(t *testing.T) {
	testCases := []struct {
		insert []int
		delete []int
		length int
	}{
		{insert: []int{5, 3, 69, 420, 4, 1, 8, 7}, delete: []int{1, 3, 4, 5, 7, 8, 69, 420}, length: 0},
		{insert: []int{5, 3, 69, 420, 4, 1, 8, 7}, delete: []int{1, 3, 4, 5, 7, 8, 69}, length: 1},
		{insert: []int{5, 3, 69, 420, 4, 1, 8, 7}, delete: []int{1, 3, 4, 5, 7, 8}, length: 2},
		{insert: []int{5, 3, 69, 420, 4, 1, 8, 7}, delete: []int{1, 3, 4, 5, 7}, length: 3},
		{insert: []int{5, 3, 69, 420, 4, 1, 8, 7}, delete: []int{1, 3, 4, 5}, length: 4},
		{insert: []int{5, 3, 69, 420, 4, 1, 8, 7}, delete: []int{1, 3, 4}, length: 5},
		{insert: []int{5, 3, 69, 420, 4, 1, 8, 7}, delete: []int{1, 3}, length: 6},
		{insert: []int{5, 3, 69, 420, 4, 1, 8, 7}, delete: []int{1}, length: 7},
		{insert: []int{5, 3, 69, 420, 4, 1, 8, 7}, delete: []int{}, length: 8},
		{insert: []int{1}, delete: []int{1}, length: 0},
		{insert: []int{}, delete: []int{-1}, length: 0},
		{insert: []int{1}, delete: []int{}, length: 1},
		{insert: []int{}, delete: []int{}, length: 0},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			heap := heap.NewMinHeap()

			for _, val := range tc.insert {
				heap.Insert(val)
			}

			for _, val := range tc.delete {
				deleted := heap.Delete()
				if deleted != val {
					t.Errorf("[case: %s] got %v, want %v", index, deleted, val)
				}
			}

			length := heap.Length()
			if length != tc.length {
				t.Errorf("[case: %s] got %v, want %v", index, length, tc.length)
			}
		})
	}
}
