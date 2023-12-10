package day10

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	y int // y is first, and goes down, 0 is the top
	x int
}

type animal struct {
	location  point
	direction string
}

func (a *animal) next(grid []string) {
	switch a.direction {
	case "N":
		a.location = point{a.location.y - 1, a.location.x}
	case "S":
		a.location = point{a.location.y + 1, a.location.x}
	case "E":
		a.location = point{a.location.y, a.location.x + 1}
	case "W":
		a.location = point{a.location.y, a.location.x - 1}
	}

	switch a.direction + string(grid[a.location.y][a.location.x]) { // "*|", "*-"': stays the same
	case "NF", "SL":
		a.direction = "E"
	case "N7", "SJ":
		a.direction = "W"
	case "WF", "E7":
		a.direction = "S"
	case "EJ", "WL":
		a.direction = "N"
	}
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	p1, p2 := 0, 0
	grid := []string{}
	start := point{}

	for y, l := range strings.Split(strings.TrimSpace(string(inpBytes)), "\n") {
		grid = append(grid, l)

		if strings.Contains(l, "S") {
			start = point{y: y, x: strings.Index(l, "S")}
		}
	}

	a := animal{location: start, direction: "E"}
	b := animal{location: start, direction: "N"}

	for {
		a.next(grid)
		b.next(grid)
		p1++

		if a.location.y == b.location.y && a.location.x == b.location.x {
			break
		}
	}

	fmt.Printf("P1: %d\nP2: %d\n", p1, p2)
}
