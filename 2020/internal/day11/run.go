package day11

import (
	"fmt"
	"os"
	"strings"
)

type coord struct {
	y int
	x int
}

func newState(y int, x int, g []string) rune {
	seatVal := g[y][x]

	if seatVal == '.' {
		return '.'
	}

	neighbors := 0

	for _, p := range []coord{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
		nY := y + p.y
		nX := x + p.x

		if nY < 0 || nY >= len(g) || nX < 0 || nX >= len(g[0]) {
			continue
		}

		if g[nY][nX] == '#' {
			neighbors++
		}
	}

	if seatVal == 'L' && neighbors == 0 {
		return '#'
	} // else

	if seatVal == '#' && neighbors >= 4 {
		return 'L'
	}

	return rune(g[y][x]) // not sure why I need to cast rune()
}

func doRound(inGrid []string) []string {
	outGrid := []string{}

	for y := range inGrid {
		row := ""

		for x := range inGrid[y] {
			n := newState(y, x, inGrid)
			row += string(n)
		}

		outGrid = append(outGrid, row)
	}

	return outGrid
}

func printGrid(g []string) {
	for y := range g {
		fmt.Printf(" %s\n", g[y])
	}
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	grid := []string{}
	grid = append(grid, strings.Split(strings.TrimSpace(inpStr), "\n")...)
	previous := strings.Join(grid, "")

	for {
		// printGrid(grid)
		newGrid := doRound(grid)
		next := strings.Join(newGrid, "")

		if previous == next {
			fmt.Printf("P1: %d\n", strings.Count(next, "#"))
			break
		}

		previous = next

		for i := range newGrid {
			grid[i] = strings.Clone(newGrid[i])
		}
	}
}
