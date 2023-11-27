package day22b

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type board struct {
	layout [][]rune
	pStat  pacman
}

type point struct{ y, x int }

type pacman struct {
	loc point
	dir rune
}

type move struct {
	howMany int
	turn    rune
}

func (b board) print() {
	fmt.Printf("y:%d, x:%d, dir:%c\n", b.pStat.loc.y, b.pStat.loc.x, b.pStat.dir)

	for y := range b.layout {
		fmt.Printf("%s\n", string(b.layout[y]))
	}
}

func parseMoveStr(in string) []move {
	moves := make([]move, 0)
	r := regexp.MustCompile(`(\d+)([^\d])`)
	matches := r.FindAllStringSubmatch(in+"X", -1) // add X for regex on last m

	for _, m := range matches {
		hm, _ := strconv.Atoi(m[1])
		dir := rune(m[2][0])
		moves = append(moves, move{howMany: hm, turn: dir})
	}

	return moves
}

// had to get a dice to visualize this. 🎲
// Smarter people will write code for this, I'm gonna be happy if I can make this hardconding work...
// A side is 50x50 (and don't forget the border of X), and the layout is:
// .   +---+---+
// .   | 1 | 2 |
// .   +---+---+
// .   | 4 |
// +---+---+      T
// | 5 | 6 |     L R
// +---+---+      B
// | 3 |
// +---+
// nolint: wsl
func transform(y, x int) (int, int, rune) {
	fmt.Printf("transform y:%d, x:%d \n", y, x)
	if x == 50 && y < 100 {
		if y <= 50 { //
			return 151 - y, 1, 'R'
		} // else
		return 100, y - 50, 'D' //
	} // else

	if x == 0 {
		if y >= 100 && y <= 150 { //
			return 150 - y, 51, 'R'
		} // else
		return 1, 250 - y, 'D' //
	} // else

	if y == 0 { // top 1 2
		if x <= 150 {
			return 200 - (x - 50), 1, 'R' //
		} // else
		return 200, x - 100, 'U' //
	} // else

	if x == 101 {
		if y <= 100 {
			return 50, y, 'U' // 3
		} // else
		return -1, -1, 'L' //
	} // else

	if y == 100 {
		return 100 - x, 50, 'R' //
	} // else

	if y == 151 {
		return x, 50, 'L' //
	} // else

	if x == 51 {
		return -1, -1, 'U' //
	} // else

	if x == 151 {
		return -1, -1, 'L' //
	}

	if y == 51 {
		return x, 100, 'L' //
	}

	if y == 201 {
		return 1, 100 + x, 'D'
	}

	panic(fmt.Sprintf("Transform with y:%d and x:%d is illegal", y, x))
}

func (b *board) move(num int) {
	indicator := map[rune]rune{'U': '^', 'D': 'v', 'R': '>', 'L': '<'}
	movesDone := 0
	lastGoodX := b.pStat.loc.x
	lastGoodY := b.pStat.loc.y
	// lastGoodDir := b.pStat.dir

	for {
		transformed := false
		nextX := b.pStat.loc.x
		nextY := b.pStat.loc.y

		switch b.pStat.dir {
		case 'U':
			nextY = (nextY - 1) % len(b.layout)
			if nextY < 0 {
				nextY = len(b.layout) - 1
			}
		case 'D':
			nextY = (nextY + 1) % len(b.layout)
		case 'R':
			nextX = (nextX + 1) % len(b.layout[0])
		case 'L':
			nextX = (nextX - 1) % len(b.layout[0])
			if nextX < 0 {
				nextX = len(b.layout[0]) - 1
			}
		default:
			panic("illegal direction")
		}

		if b.layout[nextY][nextX] == 'X' {
			// b.print()
			var dir rune

			nextY, nextX, dir = transform(nextY, nextX)
			b.pStat.dir = dir
			transformed = true
		}

		if b.layout[nextY][nextX] == '#' {
			b.pStat.loc.x = lastGoodX
			b.pStat.loc.y = lastGoodY

			break
		}

		b.layout[nextY][nextX] = indicator[b.pStat.dir]
		lastGoodX = nextX
		lastGoodY = nextY
		movesDone++

		b.pStat.loc.x = nextX
		b.pStat.loc.y = nextY

		if transformed {
			b.print()
		}

		if movesDone == num {
			break
		}
	}
}

func (b *board) turn(dir rune) {
	clockwise := map[rune]rune{'U': 'R', 'R': 'D', 'D': 'L', 'L': 'U'}
	counter := map[rune]rune{'L': 'D', 'D': 'R', 'R': 'U', 'U': 'L'}

	switch dir {
	case 'R':
		b.pStat.dir = clockwise[b.pStat.dir]
	case 'L':
		b.pStat.dir = counter[b.pStat.dir]
	default:
		fmt.Printf("illegal direction... is it the last? %q\n", dir)
	}
}

func (b *board) doMoves(moves []move) {
	for _, m := range moves {
		b.move(m.howMany)
		b.turn(m.turn)
	}
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	b := board{
		layout: [][]rune{},
		pStat: pacman{
			loc: point{y: 0, x: 0},
			dir: 'R',
		},
	}

	b.layout = append(b.layout, make([]rune, 0)) // create room for a boundary of X later
	maxX := 0

	for y, l := range strings.Split(inpStr, "\n") {
		b.layout = append(b.layout, make([]rune, 0))
		l = strings.ReplaceAll(l, " ", "X")
		b.layout[y+1] = []rune("X" + l)

		if b.pStat.loc.x <= 0 {
			b.pStat.loc.y = y + 1
			b.pStat.loc.x = strings.Index(l, ".") + 1
		}

		maxX = max(maxX, len(l))

		if l == "" {
			break
		}
	}

	moveStr := strings.TrimSuffix(strings.Split(inpStr, "\n\n")[1], "\n")
	moves := parseMoveStr(moveStr)

	for y, row := range b.layout {
		for x := len(row); x < maxX+2; x++ {
			b.layout[y] = append(b.layout[y], 'X')
		}
	}

	b.layout[b.pStat.loc.y][b.pStat.loc.x] = '>'
	b.doMoves(moves)
	b.print()

	faceVal := map[rune]int{'R': 0, 'D': 1, 'L': 2, 'U': 3}
	fmt.Printf("x: %d y: %d p1: %d \n", b.pStat.loc.x, b.pStat.loc.y, 1000*b.pStat.loc.y+4*b.pStat.loc.x+faceVal[b.pStat.dir])
}
