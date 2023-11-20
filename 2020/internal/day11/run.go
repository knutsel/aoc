package day11

import (
	"fmt"
	"os"
	"strings"
)

// type pos int

// const (
// 	floor = iota
// 	empty
// 	occupied
// )

// func runeToPos(c rune) pos {
// 	switch c {
// 	case '.':
// 		return floor
// 	case 'L':
// 		return empty
// 	case '#':
// 		return occupied
// 	default:
// 		panic("Unexpected input")
// 	}
// }

// func newState(y int, x int, grid [][]pos) pos {
// 	return -1
// }

// func doMove(grid [][]pos) ([][]pos, int) {
// 	newGrid := [][]pos{}

// 	for y := range grid {
// 		newRow := []pos{}
// 		for x := range grid[y] {
// 			newRow = append(newRow, newState(y, x, grid))
// 		}

// 		newGrid = append(newGrid, newRow)
// 	}

// 	return newGrid, -1
// }

// func inBounds(y int, x int, g []string) bool {
// 	if x < 0 || y < 0 {
// 		return false
// 	}
// 	if x >= len(g[0]) || y >= len(g) {
// 		return false
// 	}

//		return true
//	}
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

	// for y:= range []int{}
	// 	if y < 0 || y >= len(g) {
	// 		continue
	// 	}

	// 	for x := x - 1; x <= x+1; x++ {
	// 		if x < 0 || x >= len(g[0]) {
	// 			continue
	// 		}

	// 		if g[y][x] == '#' {
	// 			neighbors++
	// 		}
	// 	}
	// }

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
	// grid := [][]pos{} // y, x
	grid := []string{}

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		// row := []pos{}
		// for _, c := range l {
		// 	row = append(row, runeToPos(c))
		// }

		// grid = append(grid, row)
		grid = append(grid, l)
	}

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
