package search_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/go-dsa/search"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	testCases := []struct {
		num  int
		want bool
	}{
		{69, true},
		{1336, false},
		{69420, true},
		{69421, false},
		{1, true},
		{0, false},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := search.LinearSearch(arr, tc.num)
			if got != tc.want {
				t.Errorf("[case: %s] want %v, got %v", index, tc.want, got)
				return
			}
		})
	}
}
