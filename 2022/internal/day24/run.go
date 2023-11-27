package day24

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// note... y is first in g[y][x]
type point struct {
	y int
	x int
}

type cell struct {
	// storms   []rune
	distance int
	previous point
}

type state struct {
	pos point
	min int
}

type valley struct {
	weather    [][]rune
	start, end point
}

var (
	stormCache = map[string]bool{}
	thisValley valley
)

// going to try this with Dijkstra
// vv ----- From https://github.com/albertorestifo/dijkstra

// Queue is a basic priority queue implementation, where the node with the
// lowest priority is kept as first element in the queue
type Queue struct {
	keys  []state
	nodes map[state]int
}

// Len is part of sort.Interface
func (q *Queue) Len() int {
	return len(q.keys)
}

// Swap is part of sort.Interface
func (q *Queue) Swap(i, j int) {
	q.keys[i], q.keys[j] = q.keys[j], q.keys[i]
}

// Less is part of sort.Interface
func (q *Queue) Less(i, j int) bool {
	a := q.keys[i]
	b := q.keys[j]

	return q.nodes[a] < q.nodes[b]
}

// Set updates or inserts a new key in the priority queue
func (q *Queue) Set(key state, priority int) {
	// inserts a new key if we don't have it already
	if _, ok := q.nodes[key]; !ok {
		q.keys = append(q.keys, key)
	}

	// set the priority for the key
	q.nodes[key] = priority

	// sort the keys array
	sort.Sort(q)
}

// Next removes the first element from the queue and returns it's key and priority
func (q *Queue) Next() (key state, priority int) {
	// shift the key form the queue
	key, keys := q.keys[0], q.keys[1:]
	q.keys = keys

	priority = q.nodes[key]

	delete(q.nodes, key)

	return key, priority
}

// IsEmpty returns true when the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.keys) == 0
}

// Get returns the priority of a passed key
func (q *Queue) Get(key state) (priority int, ok bool) {
	priority, ok = q.nodes[key]
	return
}

// NewQueue creates a new empty priority queue
func NewQueue() *Queue {
	var q Queue
	q.nodes = make(map[state]int)

	return &q
}

// ^^----- From https://github.com/albertorestifo/dijkstra

// func printGrid(g [][]cell) {
// 	fmt.Println()

// 	for _, row := range g {
// 		for _, cell := range row {
// 			switch len(cell.storms) {
// 			case 0:
// 				fmt.Print(".")
// 			case 1:
// 				fmt.Printf("%c", cell.storms[0])
// 			default:
// 				fmt.Printf("%d", len(cell.storms))
// 			}
// 		}

// 		fmt.Println()
// 	}
// }

// // moveStorms moves all the storms to their next location in a new fresh (deep) copy of the grid
// func moveStorms(in [][]cell) [][]cell {
// 	out := [][]cell{}
// 	for y := range in {
// 		out = append(out, []cell{})
// 		for x := 0; x < len(in[y]); x++ {
// 			out[y] = append(out[y], cell{})
// 		}
// 	}

// 	for y, row := range in {
// 		for x, cell := range row {
// 			for _, c := range cell.storms {
// 				destX := x
// 				destY := y

// 				switch c {
// 				case '^':
// 					destY--
// 				case 'v':
// 					destY++
// 				case '<':
// 					destX--
// 				case '>':
// 					destX++
// 				}

// 				if destX == 0 && c != '#' {
// 					destX = len(row) - 2
// 				} else if destX == len(row)-1 && c != '#' {
// 					destX = 1
// 				}

// 				if destY == 0 && c != '#' {
// 					destY = len(in) - 2
// 				} else if destY == len(in)-1 && c != '#' {
// 					destY = 1
// 				}

// 				out[destY][destX].storms = append(out[destY][destX].storms, c)
// 			}
// 		}
// 	}

// 	return out
// }

// func stringify(g [][][]rune, start point) string {
// 	str := fmt.Sprintf("%d,%d", start.y, start.x)

// 	for y := 1; y < len(g)-2; y++ {
// 		for x := 1; x < len(g[0])-2; x++ {
// 			str += ":" + string(g[y][x])
// 		}
// 	}

// 	return str
// }

func (v valley) isInbounds(p point) bool {
	if (p.y == v.start.y && p.x == v.start.x) || (p.y == v.end.y && p.x == v.end.x) {
		return true
	}

	if p.x < 1 || p.y < 1 {
		return false
	}

	if p.y > len(v.weather)-2 || p.x > len(v.weather[0])-2 {
		return false
	}

	return true
}

func showPath(g [][]cell, end point) {
	next := end

	for {
		fmt.Printf("(%d,%d) @ %d \n", next.y, next.x, g[next.y][next.x].distance)
		next = g[next.y][next.x].previous

		if next.x == -1 {
			break
		}
	}
}

func (v valley) hasStorm(p point, minute int) bool {
	seenKey := fmt.Sprintf("%d,%d-%d", p.y, p.x, minute)
	if val, ok := stormCache[seenKey]; ok {
		return val
	}

	// v (there are no ^ or v in the rows with the gap in the wall
	vertY := (p.y - minute) % (len(v.weather) - 2)
	if vertY < 1 {
		vertY += len(v.weather) - 2
	}

	if v.weather[vertY][p.x] == 'v' {
		fmt.Printf("%c true for min:%d, p:%+v, -> vert:%d\n", v.weather[vertY][p.x], minute, p, vertY)

		stormCache[seenKey] = true

		return true
	}

	// ^
	vertY = (p.y + minute) % (len(v.weather) - 1)
	if v.weather[vertY][p.x] == '^' {
		fmt.Printf("%c true for min:%d, p:%+v, -> vert:%d\n", v.weather[vertY][p.x], minute, p, vertY)

		stormCache[seenKey] = true

		return true
	}

	// >
	horX := (p.x - minute) % (len(v.weather[0]) - 2)
	if horX < 1 {
		horX += len(v.weather[0]) - 2
	}

	if v.weather[p.y][horX] == '>' {
		fmt.Printf("%c true for min:%d, p:%+v, -> vert:%d\n", v.weather[p.y][horX], minute, p, vertY)

		stormCache[seenKey] = true

		return true
	}

	// <
	horX = (p.x + minute) % (len(v.weather[0]) - 1)
	if v.weather[p.y][horX] == '<' {
		fmt.Printf("%c true for min:%d, p:%+v, -> vert:%d\n", v.weather[p.y][horX], minute, p, vertY)

		stormCache[seenKey] = true

		return true
	}

	fmt.Printf("%c false for min:%d, p:%+v, -> vertY:%d\n", v.weather[vertY][p.x], minute, p, vertY)
	stormCache[seenKey] = false
	return false
}

func searchWithDijkstra(g [][]cell, start point) {
	g[start.y][start.x].distance = 0

	pq := NewQueue()

	for y, row := range g {
		for x := range row {
			if thisValley.isInbounds(point{y, x}) {
				pq.Set(state{pos: point{y, x}, min: 0}, 1e12)
			}
		}
	}

	pq.Set(state{pos: start, min: 0}, 0)

	minute := 0

	for {
		if pq.IsEmpty() {
			break
		}

		u, _ := pq.Next()

		minute++

		for _, step := range []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {0, 0}} {
			v := point{u.pos.y + step.y, u.pos.x + step.x}
			// if thisValley.isInbounds(v) {
			if thisValley.isInbounds(v) && !thisValley.hasStorm(v, minute) {
				alt := g[u.pos.y][u.pos.x].distance
				if alt < g[v.y][v.x].distance {
					g[v.y][v.x].distance = alt + minute
					g[v.y][v.x].previous = point{u.pos.y, u.pos.x}
					pq.Set(state{pos: point{v.y, v.x}, min: minute}, alt)
				}
			}
		}
	}
}

// func searchWithDijkstra(graph []vertex, start point) (map[point]int, map[point]point) {
// 	dist := make(map[point]int)
// 	prev := make(map[point]point)
// 	dist[start] = 0

// 	for _, v:= range graph {
// 		if v ==  {

// 		}
// 	}
// 	return dist, prev

// }

// func findAllVertices(g [][][]rune) []vertex {
// 	vertices := make([]vertex, 0)

// 	for y, row := range g {
// 		for x, cell := range row {
// 			if len(cell) != 0 && cell[0] != '#' {
// 				for _, step := range []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
// 					from := point{y: y, x: x}
// 					next := point{from.y + step.y, from.x + step.x}

// 					if isInbounds(g, next) {
// 						vertices = append(vertices, vertex{from: from, to: next})
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return vertices
// }

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	expeditionState := [][]cell{}
	thisValley = valley{
		weather: [][]rune{},
		start:   point{},
		end:     point{},
	}

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		row := make([]cell, 0)
		wRow := make([]rune, 0)

		for _, c := range l {
			loc := cell{
				// storms:   (make([]rune, 0)),
				distance: 1e12,
				previous: point{-1, -1},
			}
			// if c != '.' {
			// 	loc.storms = append(loc.storms, c)
			// }

			wRow = append(wRow, c)
			row = append(row, loc)
		}

		expeditionState = append(expeditionState, row)
		thisValley.weather = append(thisValley.weather, wRow)
	}

	// printGrid(currentGrid)
	// newGrid := move(currentGrid)
	// printGrid(newGrid)
	thisValley.end = point{len(expeditionState) - 1, len(expeditionState[len(expeditionState)-1]) - 2}
	thisValley.start = point{0, 1}

	// for i := 0; i < 10; i++ {
	// 	thisValley.hasStorm(point{1, 2}, i)
	// }
	searchWithDijkstra(expeditionState, point{0, 1})
	showPath(expeditionState, thisValley.end)

	for y, row := range expeditionState {
		for x, c := range row {
			if c.distance < 100000 {
				fmt.Printf("y:%d, x:%d, cell %+v\n", y, x, c)
			}
		}
	}

	fmt.Printf("%+v\n", expeditionState)

	// searchWithWeather(currentGrid, image.Point{0, 1}, 0)
	// vertices := findAllVertices(currentGrid)
	// fmt.Printf("vertices: %+v (len:%d)\n", vertices, len(vertices))
}
