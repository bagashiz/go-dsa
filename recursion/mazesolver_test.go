package recursion_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/go-dsa/recursion"
	"github.com/bagashiz/go-dsa/testutil"
)

func TestMazeSolver(t *testing.T) {
	testCases := []struct {
		maze       []string
		wall       string
		want       []recursion.Point
		start, end recursion.Point
	}{
		{
			maze: []string{
				"########## #",
				"#        # #",
				"#        # #",
				"# ######## #",
				"#          #",
				"# ##########",
			},
			wall: "#",
			want: []recursion.Point{
				{10, 0},
				{10, 1},
				{10, 2},
				{10, 3},
				{10, 4},
				{9, 4},
				{8, 4},
				{7, 4},
				{6, 4},
				{5, 4},
				{4, 4},
				{3, 4},
				{2, 4},
				{1, 4},
				{1, 5},
			},
			start: recursion.Point{X: 10, Y: 0},
			end:   recursion.Point{X: 1, Y: 5},
		},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := recursion.SolveMaze(tc.maze, tc.wall, tc.start, tc.end)

			wantedPath := drawPath(tc.maze, tc.want)
			gotPath := drawPath(tc.maze, got)

			if diff := testutil.Diff(gotPath, wantedPath); diff != "" {
				t.Errorf("[case: %s] %s %s", index, testutil.Callers(), diff)
			}
		})
	}
}

func drawPath(data []string, path []recursion.Point) []string {
	// convert data to a mutable format ([]rune slices)
	data2 := make([][]rune, len(data))
	for i, row := range data {
		data2[i] = []rune(row)
	}

	// mark '*' at each point in the path
	for _, p := range path {
		if p.Y >= 0 && p.Y < len(data2) && p.X >= 0 && p.X < len(data2[p.Y]) {
			data2[p.Y][p.X] = '*'
		}
	}

	// convert back to []string
	result := make([]string, len(data2))
	for i, row := range data2 {
		result[i] = string(row)
	}

	return result
}
