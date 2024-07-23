package recursion

/**
 * Input is array of strings that make up to look like a 2D maze.
 * Find a way from start to end point.
 * Base cases:
 * 1. It's a wall
 * 2. Off the map
 * 3. It's the end
 * 4. If already walked through
 */

type Point struct {
	X int
	Y int
}

type pointStack []Point

func (p *pointStack) push(data Point) {
	*p = append(*p, data)
}

func (p *pointStack) pop() Point {
	data := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return data
}

func SolveMaze(maze []string, wall string, start, end Point) []Point {
	seen := make([][]bool, 0)
	path := make(pointStack, 0)

	for range len(maze) {
		seen = append(seen, make([]bool, len(maze[0])))
	}

	_ = walk(maze, wall, start, end, seen, &path)

	return path
}

func walk(maze []string, wall string, current, end Point, seen [][]bool, path *pointStack) bool {
	// off the map
	if current.X < 0 || current.X >= len(maze[0]) ||
		current.Y < 0 || current.Y >= len(maze) {
		return false
	}

	// it's a wall
	if string(maze[current.Y][current.X]) == wall {
		return false
	}

	// it's the end
	if current.X == end.X && current.Y == end.Y {
		path.push(end)
		return true
	}

	// already walked through
	if seen[current.Y][current.X] {
		return false
	}

	// pre-condition
	seen[current.Y][current.X] = true
	path.push(current)

	// recurse
	dir := [][]int{
		{-1, 0}, // left
		{1, 0},  // right
		{0, -1}, // down
		{0, 1},  // up
	}

	for i := 0; i < len(dir); i++ {
		x, y := dir[i][0], dir[i][1]
		if walk(maze, wall,
			Point{current.X + x, current.Y + y},
			end, seen, path) {
			return true
		}
	}

	// post-condition
	_ = path.pop()

	return false
}
