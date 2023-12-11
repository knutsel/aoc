package day11

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	y int64 // y is first, and goes down, 0 is the top
	x int64
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}

	return x
}

func printGrid(g []string) {
	for _, l := range g {
		fmt.Printf("%s\n", l)
	}
}

func countCrossings(g []string, p1, p2 point) int64 {
	numCrossings := int64(0)

	for y := min(p1.y, p2.y); y < max(p1.y, p2.y); y++ {
		if g[y][0] == 'v' {
			numCrossings++
		}
	}

	for x := min(p1.x, p2.x); x < max(p1.x, p2.x); x++ {
		if g[0][x] == 'v' {
			numCrossings++
		}
	}

	return numCrossings
}

func sumPaths(g []string, voidSize int64) int64 {
	galaxies := map[point]bool{}

	for y := range g {
		for x := range g[y] {
			if g[y][x] == '#' {
				galaxies[point{int64(y), int64(x)}] = true
			}
		}
	}

	done := map[string]bool{}
	sum := int64(0)
	for p1, _ := range galaxies {
		for p2, _ := range galaxies {
			if done[fmt.Sprintf("%v-%v", p1, p2)] {
				continue
			}

			done[fmt.Sprintf("%v-%v", p1, p2)] = true
			done[fmt.Sprintf("%v-%v", p2, p1)] = true
			voidCrossings := countCrossings(g, p1, p2)
			dist := Abs(p1.x-p2.x) + Abs(p1.y-p2.y) + (voidSize-1)*voidCrossings
			// fmt.Printf("%+v -> %+v == %d with %d crossings \n", p1, p2, dist, voidCrossings)
			sum += dist
		}
	}

	return sum
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	grid := []string{}
	xLocWithGalaxy := map[int]bool{}
	yLocWithGalaxy := map[int]bool{}

	for y, l := range strings.Split(strings.TrimSpace(string(inpBytes)), "\n") {
		for x := range l {
			if l[x] == '#' {
				xLocWithGalaxy[x] = true
			}
		}

		if strings.Count(l, "#") > 0 {
			yLocWithGalaxy[y] = true
		}
	}

	for _, l := range strings.Split(strings.TrimSpace(string(inpBytes)), "\n") {
		expandedLine := l

		numExpansions := 0

		for x := 0; x < len(l); x++ {
			if !xLocWithGalaxy[x] {
				expandedLine = expandedLine[:x+numExpansions] + "v" + expandedLine[x+numExpansions:]
				numExpansions++
			}
		}

		grid = append(grid, expandedLine)
		if !strings.Contains(l, "#") {
			grid = append(grid, strings.Repeat("v", len(expandedLine)))
		}
	}

	// printGrid(grid)

	fmt.Printf("P1: %d\nP2: %d\n", sumPaths(grid, 1), sumPaths(grid, 999999))
}
