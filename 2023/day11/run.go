package day11

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	y int // y is first, and goes down, 0 is the top
	x int
}

// type square struct {
// 	galaxy   string
// 	location point
// }

// nolint: gochecknoglobals
// var g graph.Graph[string, Bag]

func Abs(x int) int {
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

func sumPaths(g []string) int {
	galaxies := map[point]bool{}
	for y := range g {
		for x := range g[y] {
			if g[y][x] == '#' {
				galaxies[point{y, x}] = true
			}
		}
	}

	done := map[string]bool{}
	sum := 0
	for p1, _ := range galaxies {
		for p2, _ := range galaxies {
			if done[fmt.Sprintf("%v-%v", p1, p2)] {
				continue
			}
			done[fmt.Sprintf("%v-%v", p1, p2)] = true
			done[fmt.Sprintf("%v-%v", p2, p1)] = true
			dist := Abs(p1.x-p2.x) + Abs(p1.y-p2.y)
			// fmt.Printf("%+v -> %+v == %d\n", p1, p2, dist)
			sum += dist
		}
	}

	return sum
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	// p1 := 0
	// g := graph.New(square)
	grid := []string{}
	xLocWithGalaxy := map[int]bool{}

	for _, l := range strings.Split(strings.TrimSpace(string(inpBytes)), "\n") {
		for x := range l {
			if l[x] == '#' {
				xLocWithGalaxy[x] = true
			}
		}
	}

	for _, l := range strings.Split(strings.TrimSpace(string(inpBytes)), "\n") {
		expandedLine := l

		numExpansions := 0
		for x := 0; x < len(l); x++ {
			if !xLocWithGalaxy[x] {
				expandedLine = expandedLine[:x+numExpansions] + "." + expandedLine[x+numExpansions:]
				numExpansions++
			}
		}

		grid = append(grid, expandedLine)
		if !strings.Contains(l, "#") {
			grid = append(grid, expandedLine)
		}
	}

	printGrid(grid)

	fmt.Printf("P1: %d\n", sumPaths(grid))
}
