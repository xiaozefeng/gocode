package maze

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	maze := ReadMaze("./maze.in")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

func TestWalk(t *testing.T) {
	maze := ReadMaze("./maze.in")
	steps := Walk(
		maze,
		Point{0, 0},
		Point{len(maze) - 1, len(maze[0]) - 1},
	)
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%2d ", val)
		}
		fmt.Println()
	}
}
