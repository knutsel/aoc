package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type point struct {
	y int // y is first, and goes down, 0 is the top
	x int
}

type animal struct {
	location  point
	direction string
}

func printGrid(grid []string, path, inside, outside map[point]bool) {
	for y := range grid {
		fmt.Printf("%04d  ", y)
		row := []rune(grid[y])
		for x, ch := range row {
			colorSet := 0
			p := point{y, x}
			if path[p] {
				color.Set(color.FgRed)
				colorSet++
			}
			if inside[p] {
				color.Set(color.BgCyan)
				colorSet++
			}
			if outside[p] {
				color.Set(color.FgYellow)
				colorSet++
			}
			// color.Set(color.FgRed)
			fmt.Printf("%c", ch)
			color.Unset()
			if colorSet != 1 {
				fmt.Printf("WOA at %v\n", p)
			}
		}
		fmt.Println()
	}
	// color.Cyan("Prints text in cyan.")

	// completed := 5
	// failed := 3
	// warnings := 2
	// yellow := color.New(color.FgHiWhite, color.BgYellow) // 👈

	// fmt.Printf("\n\n=== Task results ===\n\n")

	// color.Green("%d tasks completed", completed)
	// color.Red("%d tasks failed", failed)
	// yellow.Printf("%d tasks produced warnings\n", warnings) // 👈

	fmt.Println()
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
	// row := []rune(grid[a.location.y])
	// row[a.location.x] = '😃'
	// grid[a.location.y] = string(row)

}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	p1 := 0
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
	path := map[point]bool{}

	for {
		a.next(grid)
		path[a.location] = true
		b.next(grid)
		path[b.location] = true
		p1++

		if a.location.y == b.location.y && a.location.x == b.location.x {
			break
		}
	}

	// discoverd := map[point]bool{}
	// firstInside := point{start.y + 1, start.x - 3}
	// discoverd[firstInside] = true // is the explored of BFS
	// q := []point{firstInside}

	// for _, p := range q {
	// 	discoverd[p] = true

	// 	for _, neighbor := range []point{{-1, -1}, {-1, 0}, {1, 1}, {1, 0}} {
	// 		next := point{p.y + neighbor.y, p.x + neighbor.x}
	// 		fmt.Printf("next:%+v, path[next]:%+v and inside[next]:%+v\n", next, path[next], discoverd[next])
	// 		if !path[next] && !discoverd[next] {
	// 			q = append(q, next)
	// 		}
	// 	}
	// }

	inside := map[point]bool{}
	outside := map[point]bool{}

	for y := 0; y < len(grid)-1; y++ {
		for x := 0; x < len(grid[y])-1; x++ {
			p := point{y, x}

			if path[p] {
				continue
			}

			if inside[p] || outside[p] {
				continue
			}

			isInside, found := bfs(path, p)
			if isInside {
				for k, v := range found {
					inside[k] = v
				}
			} else {
				for k, v := range found {
					outside[k] = v
				}
			}
		}
	}

	printGrid(grid, path, inside, outside)
	fmt.Printf("len(path):%d, len(inside):%d, len(outside):%d\n", len(path), len(inside), len(outside))
	fmt.Printf("P1: %d\nP2: %d\n", p1, len(inside))
}

func bfs(path map[point]bool, start point) (bool, map[point]bool) {
	inside := true
	discovered := map[point]bool{}
	discovered[start] = true // is the explored of BFS
	queued := map[point]bool{}
	q := []point{start}

	for len(q) > 0 {
		p := q[0]
		discovered[p] = true
		q = q[1:]

		for _, neighbor := range []point{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
			next := point{p.y + neighbor.y, p.x + neighbor.x}
			if next.y < 0 || next.x < 0 || next.y > 139 || next.x > 139 {
				inside = false
				continue
			}
			if !path[next] && !discovered[next] && !queued[next] {
				queued[next] = true

				q = append(q, next)
			}
		}
	}
	// if inside {
	// 	inside = checkEscape(inside, path)
	// }
	return inside, discovered
}

// func checkEscape(inside, path map[point]bool) bool {
// 	for k, _ := range inside {

//		}
//	}
func main() {
	Run("input.txt")
}
