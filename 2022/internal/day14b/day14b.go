package day14b

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vertex struct {
	val rune
}

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func printGrid(g [][]vertex, min, max point) {
	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			fmt.Printf("%c", g[y][x].val)
		}
		fmt.Println()
	}
	fmt.Println()
}

func dropSand(g [][]vertex, start point, maxY int) error {
	newY := start.y + 1

	if newY == 0 {
		return errors.New("full ")
	}

	switch {
	case g[newY][start.x].val == '.': // straight down
		return dropSand(g, point{start.x, newY}, maxY)
	case g[newY][start.x-1].val == '.': // to the left & down
		return dropSand(g, point{start.x - 1, newY}, maxY)
	case g[newY][start.x+1].val == '.': // to the right & down
		return dropSand(g, point{start.x + 1, newY}, maxY)
	default:
		if start.y == 0 {
			return errors.New("full")
		}
		g[start.y][start.x].val = 'o'
		return nil
	}
}

func getPoint(in string) point {
	x, err := strconv.Atoi(strings.Split(in, ",")[0])
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(strings.Split(in, ",")[1])
	if err != nil {
		panic(err)
	}

	return point{x, y}
}

func Run() {
	fmt.Println("Starting Run (day 14b)")

	inpBytes, _ := os.ReadFile("./input.txt")
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")

	maxX := 0
	minX := int(1e9)
	maxY := maxX
	minY := minX
	pairs := make([]line, 0)

	for _, inputLine := range inputLines {
		path := strings.Split(inputLine, " -> ")
		start := getPoint(path[0])
		maxX = max(maxX, start.x)
		minX = min(minX, start.x)
		maxY = max(maxY, start.y)
		minY = min(minY, start.y)

		for i := 1; i < len(path); i++ {
			end := getPoint(path[i])
			pairs = append(pairs, line{start, end})
			maxX = max(maxX, end.x)
			minX = min(minX, end.x)
			maxY = max(maxY, end.y)
			minY = min(minY, end.y)
			start = end
		}

		// fmt.Println()
	}

	fmt.Printf("min x,y: %d,%d and max: %d,%d\n", minX, minY, maxX, maxY)

	maxY += 2
	grid := make([][]vertex, maxY+1)

	for y := 0; y < len(grid); y++ {
		grid[y] = make([]vertex, 2*(maxX+1)) // just make it 2 times bigger and 🤞
		for x := 0; x < len(grid[y]); x++ {
			grid[y][x].val = '.'
		}
	}

	pairs = append(pairs, line{point{0, maxY}, point{2 * maxX, maxY}})
	for _, pair := range pairs {
		if pair.start.x == pair.end.x { // down/up
			x := pair.start.x
			start := pair.start.y
			end := pair.end.y

			if start > end {
				start = pair.end.y
				end = pair.start.y
			}

			for y := start; y <= end; y++ {
				grid[y][x].val = 'x'
			}
		} else {
			y := pair.start.y
			start := pair.start.x
			end := pair.end.x
			if start > end {
				start = pair.end.x
				end = pair.start.x
			}
			for x := start; x <= end; x++ {
				grid[y][x].val = 'x'
			}
		}
	}

	grid[0][500].val = '+'

	number := 0

	for {
		number++

		err := dropSand(grid, point{500, 0}, maxY)
		if err != nil {
			fmt.Printf("answer: %s @ %d\n", err, number)
			break
		}

		// printGrid(grid, point{400, 0}, point{2 * maxX, maxY})
	}
}
