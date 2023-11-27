package day23

import (
	"fmt"
	"os"
	"strings"
)

var directions [4]string

type point struct {
	y, x int
}

func printGrid(c [][]rune) {
	spaceCount := 0

	for y := range c {
		fmt.Printf("%3d %s\n", y, string(c[y]))
		spaceCount += strings.Count(string(c[y]), ".")
	}

	fmt.Printf("Spaces:%d\n", spaceCount) // -80 -80 -1...
}

func expand(in [][]rune) [][]rune {
	rowLen := len(in[0])
	out := make([][]rune, 0)

	if strings.Contains(string(in[0]), "#") {
		out = append(out, []rune(strings.Repeat(".", rowLen)))
	}

	out = append(out, in...)

	if strings.Contains(string(in[len(in)-1]), "#") {
		out = append(out, []rune(strings.Repeat(".", rowLen)))
	}

	suff := ""
	pref := ""

	for i := range in {
		if strings.HasSuffix(string(in[i]), "#") {
			suff = "."
		}

		if strings.HasPrefix(string(in[i]), "#") {
			pref = "."
		}
	}

	for i := range out {
		newRow := []rune(pref + string(out[i]) + suff)
		out[i] = newRow
	}

	return out
}

func isAlone(c [][]rune, p point) bool {
	for _, n := range []point{{0, -1}, {1, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}} {
		if c[p.y+n.y][p.x+n.x] != '.' {
			return false
		}
	}

	return true
}

func propose(c [][]rune, p point, start int) point {
	for i := start; i < start+len(directions); i++ {
		dir := directions[i%(len(directions))]

		switch dir {
		case "n":
			if c[p.y-1][p.x-1] == '.' && c[p.y-1][p.x] == '.' && c[p.y-1][p.x+1] == '.' {
				return point{p.y - 1, p.x}
			}
		case "s":
			if c[p.y+1][p.x-1] == '.' && c[p.y+1][p.x] == '.' && c[p.y+1][p.x+1] == '.' {
				return point{p.y + 1, p.x}
			}
		case "w":
			if c[p.y-1][p.x-1] == '.' && c[p.y][p.x-1] == '.' && c[p.y+1][p.x-1] == '.' {
				return point{p.y, p.x - 1}
			}
		case "e":
			if c[p.y-1][p.x+1] == '.' && c[p.y][p.x+1] == '.' && c[p.y+1][p.x+1] == '.' {
				return point{p.y, p.x + 1}
			}
		}
	}

	return p
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	currentGrid := [][]rune{}
	directions = [4]string{"n", "s", "w", "e"}

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		currentGrid = append(currentGrid, []rune(l))
	}

	round := 0
	for {
		moves := make(map[point]point) // from -> to
		proposedTaken := make(map[point]bool)
		double := make(map[point]bool)

		currentGrid = expand(currentGrid)
		for y := range currentGrid {
			for x := range currentGrid[y] {
				if currentGrid[y][x] == '#' && !isAlone(currentGrid, point{y, x}) {
					newPos := propose(currentGrid, point{y, x}, round)
					moves[point{y, x}] = newPos

					if _, ok := proposedTaken[newPos]; ok {
						double[newPos] = true
					}

					proposedTaken[newPos] = true
				}
			}
		}

		if len(moves) == 0 {
			fmt.Printf("No moves at round: %d\n", round+1)
			break
		}

		for from, to := range moves {
			if _, ok := double[to]; ok {
				continue
			}

			currentGrid[from.y][from.x] = '.'
			currentGrid[to.y][to.x] = '#'
		}

		if round == 9 {
			printGrid(currentGrid)
		}
		round++
	}
}
