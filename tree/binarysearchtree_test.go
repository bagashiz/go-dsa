package tree_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/go-dsa/tree"
)

func TestBSTDepthFirstSearchInt(t *testing.T) {
	testCases := []struct {
		bst    *tree.BinaryNode[int]
		search int
		want   bool
	}{
		{bst: treesInt[0], search: 45, want: true},
		{bst: treesInt[0], search: 7, want: true},
		{bst: treesInt[0], search: 69, want: false},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := tree.BSTDepthFirstSearch(tc.bst, tc.search)
			if got != tc.want {
				t.Errorf("[case: %s] got %v, want %v", index, got, tc.want)
			}
		})
	}
}
