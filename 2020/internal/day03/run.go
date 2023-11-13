package day03

import (
	"fmt"
	"os"
	"strings"
)

func move(g [][]rune, stepX, stepY int) int {
	numTrees := 0

	for x, y := 0, 0; y < len(g); x, y = x+stepX, y+stepY {
		if g[y][x%len(g[0])] == '#' {
			numTrees++
		}
	}

	return numTrees
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	grid := [][]rune{} // it's y,x, not x,y - makes it easier to print / compare

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		row := []rune{}

		for _, c := range l {
			row = append(row, c)
		}

		grid = append(grid, row)
	}

	fmt.Printf("P1: %d\n", move(grid, 3, 1))
	fmt.Printf("P2: %d\n", move(grid, 1, 1)*move(grid, 3, 1)*move(grid, 5, 1)*move(grid, 7, 1)*move(grid, 1, 2))
}
