package tree_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/go-dsa/testutil"
	"github.com/bagashiz/go-dsa/tree"
)

func TestPreOrderTraverseInt(t *testing.T) {
	testCases := []struct {
		tree *tree.BinaryNode[int]
		want []int
	}{
		{tree: treesInt[0], want: []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}},
		{tree: treesInt[1], want: []int{20, 10, 5, 7, 15, 50, 30, 29, 21, 45, 49}},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := tree.PreOrderTraverse(tc.tree)
			if diff := testutil.Diff(got, tc.want); diff != "" {
				t.Errorf("[case: %s] got %v, want %v", index, got, tc.want)
			}
		})
	}
}
