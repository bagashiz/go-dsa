package sort_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/go-dsa/sort"
)

func TestSinglyLinkedListAppendInt(t *testing.T) {
	testCases := []struct {
		input  []int
		want   []int
		length int
	}{
		{input: []int{1, 2, 3, 4, 5}, want: []int{1, 2, 3, 4, 5}, length: 5},
		{input: []int{1}, want: []int{1}, length: 1},
		{input: []int{}, want: []int{}, length: 0},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			list := sort.NewSinglyLinkedList[int]()

			for _, v := range tc.input {
				list.Append(v)
			}

			if list.Length() != tc.length {
				t.Errorf("[case %s] Length: got %v, want %v", index, list.Length(), tc.length)
			}

			for i, want := range tc.want {
				got := list.Get(i)
				if got != want {
					t.Errorf("[case %s] Get(%v): got %v, want %v", index, i, got, want)
				}
			}
		})
	}
}

func TestSinglyLinkedListPrependInt(t *testing.T) {
	testCases := []struct {
		input  []int
		want   []int
		length int
	}{
		{input: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}, length: 5},
		{input: []int{1}, want: []int{1}, length: 1},
		{input: []int{}, want: []int{}, length: 0},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			list := sort.NewSinglyLinkedList[int]()

			for _, v := range tc.input {
				list.Prepend(v)
			}

			if list.Length() != tc.length {
				t.Errorf("[case %s] Length: got %v, want %v", index, list.Length(), tc.length)
			}

			for i, want := range tc.want {
				got := list.Get(i)
				if got != want {
					t.Errorf("[case %s] Get(%v): got %v, want %v", index, i, got, want)
				}
			}
		})
	}
}

func TestSinglyLinkedListInsertAtInt(t *testing.T) {
	type item struct {
		index int
		data  int
	}

	testCases := []struct {
		input  []item
		want   []int
		length int
	}{
		{input: []item{{index: 0, data: 1}, {index: 1, data: 2}, {index: 2, data: 3}, {index: 3, data: 4}, {index: 4, data: 5}}, want: []int{1, 2, 3, 4, 5}, length: 5},
		{input: []item{{index: 0, data: 1}, {index: 0, data: 2}, {index: 0, data: 3}, {index: 0, data: 4}, {index: 0, data: 5}}, want: []int{5, 4, 3, 2, 1}, length: 5},
		{input: []item{{index: 0, data: 1}, {index: 0, data: 2}, {index: 1, data: 3}}, want: []int{2, 3, 1}, length: 3},
		{input: []item{{index: 0, data: 1}, {index: 1, data: 2}, {index: 1, data: 3}}, want: []int{1, 3, 2}, length: 3},
		{input: []item{{index: 0, data: 1}}, want: []int{1}, length: 1},
		{input: []item{{index: -1, data: 1}}, want: []int{}, length: 0},
		{input: []item{}, want: []int{}, length: 0},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			list := sort.NewSinglyLinkedList[int]()

			for _, item := range tc.input {
				list.InsertAt(item.data, item.index)
			}

			if list.Length() != tc.length {
				t.Errorf("[case %s] Length: got %v, want %v", index, list.Length(), tc.length)
			}

			for i, want := range tc.want {
				got := list.Get(i)
				if got != want {
					t.Errorf("[case %s] Get(%v): got %v, want %v", index, i, got, want)
				}
			}
		})
	}
}

func TestSinglyLinkedListRemoveInt(t *testing.T) {
	testCases := []struct {
		append []int
		remove []int
		want   []int
		length int
	}{
		{append: []int{1, 2, 3, 4, 5}, remove: []int{1, 2, 3, 4, 5}, want: []int{}, length: 0},
		{append: []int{1, 2, 3, 4, 5}, remove: []int{1, 2, 3, 4}, want: []int{5}, length: 1},
		{append: []int{1, 2, 3, 4, 5}, remove: []int{1, 2, 3}, want: []int{4, 5}, length: 2},
		{append: []int{1, 2, 3, 4, 5}, remove: []int{1, 2}, want: []int{3, 4, 5}, length: 3},
		{append: []int{1, 2, 3, 4, 5}, remove: []int{1}, want: []int{2, 3, 4, 5}, length: 4},
		{append: []int{1, 2, 3, 4, 5}, remove: []int{6}, want: []int{1, 2, 3, 4, 5}, length: 5},
		{append: []int{1, 2, 3, 4, 5}, remove: []int{}, want: []int{1, 2, 3, 4, 5}, length: 5},
		{append: []int{}, remove: []int{1}, want: []int{}, length: 0},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			list := sort.NewSinglyLinkedList[int]()

			for _, v := range tc.append {
				list.Append(v)
			}

			for _, v := range tc.remove {
				list.Remove(v)
			}

			if list.Length() != tc.length {
				t.Errorf("[case %s] Length: got %v, want %v", index, list.Length(), tc.length)
			}

			for i, want := range tc.want {
				got := list.Get(i)
				if got != want {
					t.Errorf("[case %s] Get(%v): got %v, want %v", index, i, got, want)
				}
			}
		})
	}
}

func TestSinglyLinkedListRemoveAtInt(t *testing.T) {
	testCases := []struct {
		append   []int
		removeAt []int
		want     []int
		length   int
	}{
		{append: []int{1, 2, 3, 4, 5}, removeAt: []int{4, 3, 2, 1, 0}, want: []int{}, length: 0},
		{append: []int{1, 2, 3, 4, 5}, removeAt: []int{4, 3, 2, 1}, want: []int{1}, length: 1},
		{append: []int{1, 2, 3, 4, 5}, removeAt: []int{4, 3, 2}, want: []int{1, 2}, length: 2},
		{append: []int{1, 2, 3, 4, 5}, removeAt: []int{1, 3}, want: []int{1, 3, 4}, length: 3},
		{append: []int{1, 2, 3, 4, 5}, removeAt: []int{0, 1}, want: []int{2, 4, 5}, length: 3},
		{append: []int{1, 2, 3, 4, 5}, removeAt: []int{6}, want: []int{1, 2, 3, 4, 5}, length: 5},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			list := sort.NewSinglyLinkedList[int]()

			for _, v := range tc.append {
				list.Append(v)
			}

			for _, v := range tc.removeAt {
				list.RemoveAt(v)
			}

			if list.Length() != tc.length {
				t.Errorf("[case %s] Length: got %v, want %v", index, list.Length(), tc.length)
			}

			for i, want := range tc.want {
				got := list.Get(i)
				if got != want {
					t.Errorf("[case %s] Get(%v): got %v, want %v", index, i, got, want)
				}
			}
		})
	}
}

func TestSinglyLinkedListGetInt(t *testing.T) {
	testCases := []struct {
		append []int
		index  int
		want   int
	}{
		{append: []int{1, 2, 3}, index: 0, want: 1},
		{append: []int{1, 2, 3}, index: 1, want: 2},
		{append: []int{1, 2, 3}, index: 2, want: 3},
		{append: []int{1, 2, 3}, index: 3, want: 0},
		{append: []int{1, 2, 3}, index: -1, want: 0},
		{append: []int{}, index: 0, want: 0},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			list := sort.NewSinglyLinkedList[int]()

			for _, v := range tc.append {
				list.Append(v)
			}

			got := list.Get(i)
			if got != tc.want {
				t.Errorf("[case %s] Get(%v): got %v, want %v", index, i, got, tc.want)
			}
		})
	}
}
