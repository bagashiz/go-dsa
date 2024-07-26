package tree_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/go-dsa/testutil"
	"github.com/bagashiz/go-dsa/tree"
)

func TestPostOrderTraverseInt(t *testing.T) {
	testCases := []struct {
		tree *tree.BinaryNode[int]
		want []int
	}{
		{tree: treesInt[0], want: []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}},
		{tree: treesInt[1], want: []int{7, 5, 15, 10, 21, 29, 49, 45, 30, 50, 20}},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := tree.PostOrderTraverse(tc.tree)
			if diff := testutil.Diff(got, tc.want); diff != "" {
				t.Errorf("[case: %s] got %v, want %v", index, got, tc.want)
			}
		})
	}
}
