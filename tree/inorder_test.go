package tree_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/go-dsa/testutil"
	"github.com/bagashiz/go-dsa/tree"
)

func TestInOrderTraverseInt(t *testing.T) {
	testCases := []struct {
		tree *tree.BinaryNode[int]
		want []int
	}{
		{tree: treesInt[0], want: []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}},
		{tree: treesInt[1], want: []int{5, 7, 10, 15, 20, 21, 29, 30, 45, 49, 50}},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := tree.InOrderTraverse(tc.tree)
			if diff := testutil.Diff(got, tc.want); diff != "" {
				t.Errorf("[case: %s] got %v, want %v", index, got, tc.want)
			}
		})
	}
}
