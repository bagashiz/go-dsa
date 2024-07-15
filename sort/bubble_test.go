package sort_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bagashiz/go-dsa/sort"
)

func TestBubbleSort(t *testing.T) {
	testCases := []struct {
		arr  []int
		want []int
	}{
		{arr: []int{9, 3, 7, 4, 69, 420, 42}, want: []int{3, 4, 7, 9, 42, 69, 420}},
		{arr: []int{1, 2, 3, 4, 5}, want: []int{1, 2, 3, 4, 5}},
		{arr: []int{1, 3, 2, 5, 4}, want: []int{1, 2, 3, 4, 5}},
		{arr: []int{7, 2, 9, 1, 5}, want: []int{1, 2, 5, 7, 9}},
		{arr: []int{50, 40, 30, 20, 10}, want: []int{10, 20, 30, 40, 50}},
		{arr: []int{15, 10, 25, 20, 30}, want: []int{10, 15, 20, 25, 30}},
		{arr: []int{100, 50, 200, 150, 250}, want: []int{50, 100, 150, 200, 250}},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			sort.BubbleSort(tc.arr)
			if !reflect.DeepEqual(tc.arr, tc.want) {
				t.Errorf("[case: %s] want %v, got %v", index, tc.want, tc.arr)
				return
			}
		})
	}
}
