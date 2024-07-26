package tree_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/go-dsa/tree"
)

func TestCompareBinaryTrees(t *testing.T) {
	tree3 := &tree.BinaryNode[int]{Data: 69420}

	testCases := []struct {
		tree1 *tree.BinaryNode[int]
		tree2 *tree.BinaryNode[int]
		want  bool
	}{
		{tree1: treesInt[0], tree2: treesInt[0], want: true},
		{tree1: treesInt[0], tree2: treesInt[1], want: false},
		{tree1: treesInt[0], tree2: tree3, want: false},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := tree.CompareBT(tc.tree1, tc.tree2)
			if got != tc.want {
				t.Errorf("[case: %s] got %v, want %v", index, got, tc.want)
			}
		})
	}
}
