package day12b

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
	fmt.Println("Starting Run (day 12b)")

	inpBytes, _ := os.ReadFile("./input.txt")
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")
	grid := make([][]vertex, 0)
	end := point{}
	lno := 0
	starts := make([]point, 0)

	for _, line := range inputLines {
		grid = append(grid, make([]vertex, 0))
		stat := ' '

		for cno, c := range line {
			switch c {
			case 'S':
				c = 'a'

				starts = append(starts, point{lno, cno})

			case 'E':
				c = 'z'
				end = point{lno, cno}
			case 'a':
				starts = append(starts, point{lno, cno})
			}

			grid[lno] = append(grid[lno], vertex{val: byte(c), status: byte(stat)})
		}
		lno++
	}

	results := make(map[point]int)

	for _, s := range starts {
		r := bfs(grid, s, end)
		if r > 0 {
			results[s] = r
		}

		for x, l := range grid { // reset the status'... this can be smarter
			for y := range l {
				grid[x][y].status = ' '
			}
		}
	}

	min := int(2e9)
	for _, v := range results {
		if v != -1 && v < min {
			fmt.Printf("r: %d\n", v)
			min = v
		}
	}
	fmt.Printf("----- steps: %d\n", min)
}
