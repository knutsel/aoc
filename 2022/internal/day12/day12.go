package day12

import (
	"fmt"
	"os"
	"strings"
)

type vertex struct {
	val    byte
	status byte
	dist   int
}

type point struct {
	x int
	y int
}

func printGrid(g [][]vertex) {
	for _, l := range g {
		for _, v := range l {
			fmt.Printf("%c", v.val)
		}
		fmt.Println()
	}
	fmt.Println()
	for _, l := range g {
		for _, v := range l {
			fmt.Printf("%c", v.status)
		}
		fmt.Println()
	}
}

// func dfs(grid [][]vertex, startx, starty int) {
// 	// printGrid(grid)
// 	steps++
// 	highest := grid[startx][starty].val
// 	grid[startx][starty].status = '.'
// 	// for x, y:= range []int, int[{startx +1, starty}, {startx-1, starty}]
// 	for _, x := range []int{startx, startx - 1, startx + 1} {
// 		for _, y := range []int{starty, starty - 1, starty + 1} {
// 			if x >= 0 && y >= 0 && // boundary
// 				y < len(grid[0]) && x < len(grid) && // boundary
// 				grid[x][y].status == ' ' && // already been there
// 				!(x != startx && y != starty) && // diagonally
// 				grid[x][y].val >= grid[startx][starty].val && // is it going up?
// 				grid[x][y].val-grid[startx][starty].val <= 1 { // by no more than one step
// 				fmt.Printf("start: %d,%d next: %d,%d steps: %d\n", startx, starty, x, y, steps)
// 				// if x>startx {grid[startx, starty]}
// 				dfs(grid, x, y)
// 			}
// 		}
// 	}
// 	fmt.Printf("Steps now is %d highest is %c\n", steps, highest)
// }

func bfs(grid [][]vertex, start, end point) int {
	// path:=make([]point)
	grid[start.x][start.y].status = '.'
	grid[start.x][start.y].dist = 0
	queue := []point{{start.x, start.y}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if p.x == end.x && p.y == end.y {
			return grid[p.x][p.y].dist
		}
		for _, x := range []int{p.x, p.x - 1, p.x + 1} {
			for _, y := range []int{p.y, p.y - 1, p.y + 1} {

				if x >= 0 && y >= 0 && // boundary
					y < len(grid[0]) && x < len(grid) && // boundary
					!(x != p.x && y != p.y) && // diagonally
					(grid[x][y].val >= grid[p.x][p.y].val && grid[x][y].val-grid[p.x][p.y].val <= 1 || //  going up, by no more than one step
						grid[x][y].val < grid[p.x][p.y].val) && // _or_ going down
					grid[x][y].status == ' ' { // hasn't been visited
					grid[x][y].status = grid[x][y].val
					grid[x][y].dist = grid[p.x][p.y].dist + 1
					queue = append(queue, point{x, y})
				}
			}
		}
	}
	return -1
}

func Run() {
	fmt.Println("Starting Run (day 12)")
	inpBytes, _ := os.ReadFile("./input.txt")
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")

	grid := make([][]vertex, 0)
	var start, end point
	lno := 0
	for _, line := range inputLines {
		grid = append(grid, make([]vertex, 0))
		for cno, c := range line {
			if c == 'S' {
				start = point{lno, cno}
				c = 'a'
			}
			stat := ' '
			if c == 'E' {
				c = 'z'
				end = point{lno, cno}
			}
			grid[lno] = append(grid[lno], vertex{val: byte(c), status: byte(stat)})
		}
		lno++
	}
	steps := bfs(grid, start, end)
	printGrid(grid)
	fmt.Printf("Starting point is %+v, end is %+v grid is %dx%d\n", start, end, len(grid), len(grid[0]))
	fmt.Printf("----- steps: %d\n", steps)
}
