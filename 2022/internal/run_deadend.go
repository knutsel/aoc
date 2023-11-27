package day24

import (
	"container/heap"
	"fmt"
	"image"
	"os"
	"strings"
)

// nolint: gochecknoglobals
var (
	end   image.Point
	level string
)

// vv ----- From https://pkg.go.dev/container/heap -----
// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// ^^----- From https://pkg.go.dev/container/heap -----

func printGrid(g [][][]rune) {
	fmt.Println()

	for _, row := range g {
		for _, col := range row {
			switch len(col) {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Printf("%c", col[0])
			default:
				fmt.Printf("%d", len(col))
			}
		}

		fmt.Println()
	}
}

// moveStorms all the storms to their next location in a new fresh (deep) copy of the grid
func moveStorms(in [][][]rune) [][][]rune {
	out := [][][]rune{}
	for y := range in {
		out = append(out, [][]rune{})
		for x := 0; x < len(in[y]); x++ {
			out[y] = append(out[y], []rune{})
		}
	}

	for y, row := range in {
		for x, locList := range row {
			for _, c := range locList {
				destX := x
				destY := y

				switch c {
				case '^':
					destY--
				case 'v':
					destY++
				case '<':
					destX--
				case '>':
					destX++
				}

				if destX == 0 && c != '#' {
					destX = len(row) - 2
				} else if destX == len(row)-1 && c != '#' {
					destX = 1
				}

				if destY == 0 && c != '#' {
					destY = len(in) - 2
				} else if destY == len(in)-1 && c != '#' {
					destY = 1
				}

				out[destY][destX] = append(out[destY][destX], c)
				// out[y][x] = append(out[y][x], '#')
			}
		}
	}

	return out
}

func stringify(g [][][]rune, start image.Point) string {
	str := fmt.Sprintf("%d,%d", start.Y, start.X)

	for y := 1; y < len(g)-2; y++ {
		for x := 1; x < len(g[0])-2; x++ {
			str += ":" + string(g[y][x])
		}
	}

	return str
}

func isInbounds(g [][][]rune, p image.Point) bool {
	if p.Y < len(g) && p.X < len(g[0]) && p.Y > 0 && p.X > 0 { // no -1 bc of '#' border
		return true
	}

	return false
}

func searchWithDijkstra(g [][][]rune, start, end image.Point) {
	dist := make(map[image.Point]int)
	prev := make(map[image.Point]image.Point)
	dist[start] = 0

	pq := make(PriorityQueue, 0)
	for _, step := range []image.Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		next := start.Add(step)
		// if isInbounds(newGrid, next) && len(newGrid[next.Y][next.X]) == 0 {
		dist[next] = 1e12
		prev[next] = image.Point{-1, -1}
		pq.Push(next)
		// }
	}
}
func searchWithWeather(g [][][]rune, start image.Point, minutes int) int {
	cacheKey := stringify(g, start)
	fmt.Printf("%s Start: %+v minutes:%d stringyfied:%s \n", level, start, minutes, cacheKey)
	level += "  "
	if minutes >= 30 {
		return minutes
	}
	if start == end {
		fmt.Printf("END at %d\n", minutes)
		return minutes
	}
	newGrid := moveStorms(g)
	nextMoves := make([]image.Point, 0)
	// nextMoves = append(nextMoves, start)

	for _, step := range []image.Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} { // these are sorted to go down, right first
		// nextY := start.y + dirs.y
		// nextX := start.x + dirs.x
		next := start.Add(step)
		if isInbounds(newGrid, next) && len(newGrid[next.Y][next.X]) == 0 {
			nextMoves = append(nextMoves, image.Point{next.Y, next.X})
		}
	}

	if len(nextMoves) == 0 {
		nextMoves = append(nextMoves, start)
	}

	// fmt.Printf("len moves: %d\n", len(nextMoves))
	minutes++
	for _, m := range nextMoves {
		// newGrid[start.y][start.x] = append(newGrid[start.y][start.x], '#') //?
		searchWithWeather(newGrid, m, minutes)
	}

	fmt.Printf("HERE %d - cacheKey:%s\n", minutes, cacheKey)
	return -1
}

func Run(fName string) {
	level = ""
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	currentGrid := [][][]rune{}

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		row := make([][]rune, 0)

		for _, c := range l {
			loc := (make([]rune, 0))
			if c != '.' {
				loc = append(loc, c)
			}

			row = append(row, loc)
		}

		currentGrid = append(currentGrid, row)
	}

	printGrid(currentGrid)
	// newGrid := move(currentGrid)
	// printGrid(newGrid)
	end = image.Point{len(currentGrid) - 1, len(currentGrid[len(currentGrid)-1]) - 2}
	searchWithWeather(currentGrid, image.Point{0, 1}, 0)
}
