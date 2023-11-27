package day9b

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
	maxX          int
	maxY          int
	tail          point
	head          point
	start         point
	visitedByTail map[point]bool
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (gr *grid) tailFollow() {
	dist := point{gr.head.x - gr.tail.x, gr.head.y - gr.tail.y}
	if dist.x != 0 && dist.y != 0 && (Abs(dist.x) > 1 || Abs(dist.y) > 1) {
		if dist.x > 0 {
			gr.tail.x++
		} else {
			gr.tail.x--
		}

		if dist.y > 0 {
			gr.tail.y++
		} else {
			gr.tail.y--
		}
		return
	}
	// else
	switch dist.x {
	case 2:
		gr.tail.x++
	case -2:
		gr.tail.x--
	}

	switch dist.y {
	case 2:
		gr.tail.y++
	case -2:
		gr.tail.y--
	}
}

func (gr *grid) step(dir rune) {
	// fmt.Printf("Stepping %q;", dir)
	switch dir {
	case 'U':
		gr.head.y++
	case 'D':
		gr.head.y--
	case 'R':
		gr.head.x++
	case 'L':
		gr.head.x--
	default:
		fmt.Printf("%q invalid direction\n", dir)
	}
	gr.tailFollow()
	gr.visitedByTail[gr.tail] = true
}

func (gr *grid) move(where rune, steps int) {
	// fmt.Printf("Going %q by %d: ", where, steps)
	for i := 0; i < steps; i++ {
		gr.step(where)
	}
	// fmt.Println()
}

func (gr *grid) printGrid() {
	fmt.Println("-----------------")
	for y := gr.maxY - 1; y >= -16; y-- {
		for x := -16; x < gr.maxX; x++ {
			letter := "."
			if _, ok := gr.visitedByTail[point{x, y}]; ok {
				letter = "X"
			}
			if gr.start.x == x && gr.start.y == y {
				letter = "s"
			}
			if gr.tail.x == x && gr.tail.y == y {
				letter = "T"
			}
			if gr.head.x == x && gr.head.y == y {
				letter = "H"
			}

			fmt.Printf("%s", letter)
		}
		fmt.Println()
	}
}

func Run() {
	fmt.Println("Starting Run (day 9)")
	inpBytes, _ := os.ReadFile("input.txt")
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")
	gr := grid{
		maxX:          16,
		maxY:          16,
		tail:          point{x: 0, y: 0},
		head:          point{x: 0, y: 0},
		start:         point{x: 0, y: 0},
		visitedByTail: make(map[point]bool),
	}
	// gr.printGrid()

	for _, line := range inputLines {
		var dir rune
		var steps int
		fmt.Sscanf(line, "%c %d", &dir, &steps)
		gr.move(dir, steps)

		// gr.printGrid()
	}
	fmt.Printf("tailVisited number: %d\n", len(gr.visitedByTail))
}
