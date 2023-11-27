package day22

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type board struct {
	layout       [][]rune
	pacmanStatus pacman
}

type pacman struct {
	y, x      int
	direction rune
}

type move struct {
	howMany int
	turn    rune
}

func (b board) print() {
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

func (b *board) move(num int) {
	indicator := map[rune]rune{'U': '^', 'D': 'v', 'R': '>', 'L': '<'}
	movesDone := 0
	lastGoodX := b.pacmanStatus.x
	lastGoodY := b.pacmanStatus.y

	for {
		nextX := b.pacmanStatus.x
		nextY := b.pacmanStatus.y

		switch b.pacmanStatus.direction {
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

		if b.layout[nextY][nextX] == '#' {
			b.pacmanStatus.x = lastGoodX
			b.pacmanStatus.y = lastGoodY

			break
		}

		if b.layout[nextY][nextX] != 'X' {
			b.layout[nextY][nextX] = indicator[b.pacmanStatus.direction]
			lastGoodX = nextX
			lastGoodY = nextY
			movesDone++
		}

		b.pacmanStatus.x = nextX
		b.pacmanStatus.y = nextY

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
		b.pacmanStatus.direction = clockwise[b.pacmanStatus.direction]
	case 'L':
		b.pacmanStatus.direction = counter[b.pacmanStatus.direction]
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
		pacmanStatus: pacman{
			y:         0,
			x:         0,
			direction: 'R',
		},
	}

	b.layout = append(b.layout, make([]rune, 0)) // create room for a boundary of X later
	maxX := 0

	for y, l := range strings.Split(inpStr, "\n") {
		b.layout = append(b.layout, make([]rune, 0))
		l = strings.ReplaceAll(l, " ", "X")
		b.layout[y+1] = []rune("X" + l)

		if b.pacmanStatus.x <= 0 {
			b.pacmanStatus.y = y + 1
			b.pacmanStatus.x = strings.Index(l, ".") + 1
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

	b.layout[b.pacmanStatus.y][b.pacmanStatus.x] = '>'
	b.doMoves(moves)
	b.print()

	faceVal := map[rune]int{'R': 0, 'D': 1, 'L': 2, 'U': 3}
	fmt.Printf("x: %d y: %d p1: %d \n", b.pacmanStatus.x, b.pacmanStatus.y, 1000*b.pacmanStatus.y+4*b.pacmanStatus.x+faceVal[b.pacmanStatus.direction])
}
