package day11

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	y int
	x int
}

type stateFunc func(y int, x int, g []string) rune

func newStateP2(y int, x int, g []string) rune {
	seatVal := g[y][x]

	if seatVal == '.' {
		return '.'
	}

	neighbors := 0

	for _, p := range []point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
		for mult := 1; mult < len(g); mult++ {
			nY := y + p.y*mult
			nX := x + p.x*mult

			if nY < 0 || nY >= len(g) || nX < 0 || nX >= len(g[0]) {
				break
			}

			if g[nY][nX] == 'L' {
				break
			}

			if g[nY][nX] == '#' {
				neighbors++
				break
			}
		}
	}

	if seatVal == 'L' && neighbors == 0 {
		return '#'
	} // else

	if seatVal == '#' && neighbors >= 5 {
		return 'L'
	}

	return rune(g[y][x]) // not sure why I need to cast rune()
}

func newStateP1(y int, x int, g []string) rune {
	seatVal := g[y][x]

	if seatVal == '.' {
		return '.'
	}

	neighbors := 0

	for _, p := range []point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
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

func doRound(inGrid []string, sFunc stateFunc) []string {
	outGrid := []string{}

	for y := range inGrid {
		row := ""

		for x := range inGrid[y] {
			n := sFunc(y, x, inGrid)
			row += string(n)
		}

		outGrid = append(outGrid, row)
	}

	return outGrid
}

func doRounds(grid []string, sFunc stateFunc) {
	previous := strings.Join(grid, "")

	for {
		newGrid := doRound(grid, sFunc)
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

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	grid := []string{}
	grid = append(grid, strings.Split(strings.TrimSpace(inpStr), "\n")...)

	doRounds(grid, newStateP1)

	gridP2 := []string{}
	gridP2 = append(gridP2, strings.Split(strings.TrimSpace(inpStr), "\n")...)

	doRounds(gridP2, newStateP2)
}
