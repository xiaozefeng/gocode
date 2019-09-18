package maze

import (
	"fmt"
	"os"
)

type Point struct {
	i, j int
}

func (p Point) Add(r Point) Point {
	return Point{p.i + r.i, p.j + r.j}
}

func (p Point) At(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

var dirs = [4]Point{
	{-1, 0}, // 上
	{0, -1}, // 左
	{1, 0},  // 下
	{0, 1},  // 右
}

func Walk(maze [][]int, start, end Point) [][]int {
	var steps = make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[0]))
	}

	Q := []Point{start}
	// 当队列不为空才继续
	for len(Q) > 0 {
		var cur = Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.Add(dir)

			// maze at next is 0
			// steps at next is 0
			// next != start
			if val, ok := next.At(maze); !ok || val == 1 {
				continue
			}

			if val, ok := next.At(steps); !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			var curStep, _ = cur.At(steps)
			steps[next.i][next.j] = curStep + 1

			Q = append(Q, next)
		}
	}
	return steps
}

func ReadMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	_, _ = fmt.Fscanf(file, "%d %d", &row, &col)
	// row
	result := make([][]int, row)
	// col
	for i := range result {
		result[i] = make([]int, col)
		for j := range result[i] {
			_, _ = fmt.Fscanf(file, "%d", &result[i][j])
		}
	}
	return result
}
