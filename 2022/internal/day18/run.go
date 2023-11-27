package day18

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y, z int
}

func newGrid(x, y, z int) [][][]bool {
	g := make([][][]bool, x)
	for xx := 0; xx < x; xx++ {
		g[xx] = make([][]bool, y)
		for yy := 0; yy < y; yy++ {
			g[xx][yy] = make([]bool, z)
		}
	}

	return g
}

func getsAirDfs(x, y, z int, g [][][]bool, visited map[point]bool) bool {
	if x == 0 || x == len(g)-1 || y == 0 || y == len(g[0])-1 || z == 0 || z == len(g[0][0])-1 {
		return true // at the edge of the grid
	}

	// if x == 21 && y == 9 && z == 10 {
	// 	fmt.Printf("boo")
	// }

	visited[point{x, y, z}] = true

	for _, p := range []point{{-1, 0, 0}, {+1, 0, 0}, {0, +1, 0}, {0, -1, 0}, {0, 0, +1}, {0, 0, -1}} {
		xx := x + p.x
		yy := y + p.y
		zz := z + p.z

		_, ok := visited[point{xx, yy, zz}]
		if !ok && !g[xx][yy][zz] {
			return getsAirDfs(xx, yy, zz, g, visited)
		}
	}

	// fmt.Printf("%d,%d,%d gets no air!\n", x, y, z)
	return false
}

func exposed(x, y, z int, g [][][]bool, excludePockets bool) int {
	sides := 6

	for _, p := range []point{{-1, 0, 0}, {+1, 0, 0}, {0, +1, 0}, {0, -1, 0}, {0, 0, +1}, {0, 0, -1}} {
		if g[x+p.x][y+p.y][z+p.z] || (excludePockets && !getsAirDfs(x+p.x, y+p.y, z+p.z, g, make(map[point]bool))) {
			sides--
		}
	}

	return sides
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := strings.TrimSpace(string(inpBytes))

	input := make([]point, 0)
	maxX, maxY, maxZ := 0, 0, 0

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		var x, y, z int

		fmt.Sscanf(l, "%d,%d,%d", &x, &y, &z)
		input = append(input, point{x, y, z})
		maxX = max(x, maxX)
		maxY = max(y, maxY)
		maxZ = max(z, maxZ)
	}

	grid := newGrid(maxX+3, maxY+3, maxZ+3) // +3 makes an outter layer of false around everything to work around boundary issues
	for _, kube := range input {
		grid[kube.x+1][kube.y+1][kube.z+1] = true // input begins at 0, do +1 for outter layer
	}

	totalExposed := 0
	totalExposedExclPockets := 0

	for x := 1; x < len(grid); x++ {
		for y := 1; y < len(grid[x]); y++ {
			for z := 1; z < len(grid[x][y]); z++ {
				if grid[x][y][z] {
					totalExposed += exposed(x, y, z, grid, false)
					totalExposedExclPockets += exposed(x, y, z, grid, true)
				}
			}
		}
	}

	fmt.Printf("file:%s part1: %d part2:%d \n", fName, totalExposed, totalExposedExclPockets)
}
