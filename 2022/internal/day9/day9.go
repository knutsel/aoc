package day9

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x int
	y int
}

type grid struct {
	rope          []point
	start         point
	visitedByTail map[point]bool
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func follow(first, second point) point {
	dist := point{first.x - second.x, first.y - second.y}
	if dist.x != 0 && dist.y != 0 && (Abs(dist.x) > 1 || Abs(dist.y) > 1) {
		if dist.x > 0 {
			second.x++
		} else {
			second.x--
		}

		if dist.y > 0 {
			second.y++
		} else {
			second.y--
		}
		return second
	}
	// else
	switch dist.x {
	case 2:
		second.x++
	case -2:
		second.x--
	}

	switch dist.y {
	case 2:
		second.y++
	case -2:
		second.y--
	}
	return second
}

func (gr *grid) step(dir rune) {
	// fmt.Printf("Stepping %q;", dir)
	switch dir {
	case 'U':
		gr.rope[0].y++
	case 'D':
		gr.rope[0].y--
	case 'R':
		gr.rope[0].x++
	case 'L':
		gr.rope[0].x--
	default:
		fmt.Printf("%q invalid direction\n", dir)
	}
	for i := 0; i < 9; i++ {
		gr.rope[i+1] = follow(gr.rope[i], gr.rope[i+1])
	}
	gr.visitedByTail[gr.rope[9]] = true
}

func (gr *grid) move(where rune, steps int) {
	for i := 0; i < steps; i++ {
		gr.step(where)
	}
}

func Run(fName string) {
	fmt.Printf("Starting Run (day 9 %s)\n", fName)
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")
	gr := grid{
		rope:          make([]point, 10),
		start:         point{x: 0, y: 0},
		visitedByTail: make(map[point]bool),
	}

	for _, line := range inputLines {
		var dir rune
		var steps int
		fmt.Sscanf(line, "%c %d", &dir, &steps)
		gr.move(dir, steps)
	}
	fmt.Printf("tailVisited number: %d\n", len(gr.visitedByTail))
}
