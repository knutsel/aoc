package day12

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	y int
	x int
}
type direction int

const (
	east direction = iota
	south
	west
	north
)

type ferry struct {
	location         point
	waypointLocation point
	facing           direction
}

func (f *ferry) rotate(value int) {
	howMuch := value / 90
	newVal := f.facing + direction(howMuch)

	if newVal >= 0 {
		f.facing = newVal % 4
	} else {
		f.facing = 4 + newVal
	}
}

func (f *ferry) move(m string) {
	action, value := 'x', 0

	_, err := fmt.Sscanf(m, "%c%d", &action, &value)
	if err != nil {
		panic(err)
	}

	switch action {
	case 'N':
		f.location.y += value
	case 'S':
		f.location.y -= value
	case 'E':
		f.location.x += value
	case 'W':
		f.location.x -= value
	case 'F':
		switch f.facing {
		case north:
			f.location.y += value
		case south:
			f.location.y -= value
		case east:
			f.location.x += value
		case west:
			f.location.x -= value
		}
	case 'R':
		f.rotate(value)
	case 'L':
		f.rotate(-value)
	}
}

func (f *ferry) rotateP2(value int) {
	if value < 0 {
		switch value {
		case -90:
			value = 270
		case -180:
			value = 180
		case -270:
			value = 90
		}
	}

	howMuch := value / 90

	for i := 0; i < howMuch; i++ {
		x := f.waypointLocation.y
		y := -f.waypointLocation.x
		f.waypointLocation.y = y
		f.waypointLocation.x = x
	}
}

func (f *ferry) moveP2(m string) {
	action, value := 'x', 0

	_, err := fmt.Sscanf(m, "%c%d", &action, &value)
	if err != nil {
		panic(err)
	}

	switch action {
	case 'N':
		f.waypointLocation.y += value
	case 'S':
		f.waypointLocation.y -= value
	case 'E':
		f.waypointLocation.x += value
	case 'W':
		f.waypointLocation.x -= value
	case 'F':
		xMove := (f.waypointLocation.x) * value
		yMove := (f.waypointLocation.y) * value
		f.location.x += xMove
		f.location.y += yMove
	case 'R':
		f.rotateP2(value)
	case 'L':
		f.rotateP2(-value)
	}
}

func abs(inp int) int {
	if inp < 0 {
		return -inp
	}

	return inp
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	moves := []string{}
	moves = append(moves, strings.Split(strings.TrimSpace(inpStr), "\n")...)

	f := ferry{
		location: point{0, 0},
		facing:   east,
	}

	for _, m := range moves {
		f.move(m)
	}

	fmt.Printf("P1:%d\n", abs(f.location.x)+abs(f.location.y))
	f = ferry{
		location:         point{0, 0},
		facing:           east,
		waypointLocation: point{1, 10},
	}

	for _, m := range moves {
		f.moveP2(m)
	}

	fmt.Printf("P2:%d\n", abs(f.location.x)+abs(f.location.y))
}
