package search_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/bagashiz/go-dsa/search"
)

func TestTwoCrystalBalls(t *testing.T) {
	breakpoint := rand.IntN(10000)
	levels := make([]bool, 10000)

	for i := breakpoint; i < len(levels); i++ {
		levels[i] = true
	}

	testCases := []struct {
		arr  []bool
		want int
	}{
		{arr: levels, want: breakpoint},
		{arr: make([]bool, 821), want: -1},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := search.TwoCrystalBalls(tc.arr)
			if got != tc.want {
				t.Errorf("[case: %s] want %v, got %v", index, tc.want, got)
				return
			}
		})
	}
}
