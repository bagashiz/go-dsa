package tree_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/go-dsa/tree"
)

func TestBreadthFirstSearchInt(t *testing.T) {
	testCases := []struct {
		tree   *tree.BinaryNode[int]
		search int
		want   bool
	}{
		{tree: treesInt[0], search: 45, want: true},
		{tree: treesInt[0], search: 7, want: true},
		{tree: treesInt[0], search: 69, want: false},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := tree.BreadthFirstSearch(tc.tree, tc.search)
			if got != tc.want {
				t.Errorf("[case: %s] got %v, want %v", index, got, tc.want)
			}
		})
	}
}
