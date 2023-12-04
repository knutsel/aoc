package day03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type point struct {
	y int
	x int
}

func sumParts2(s []string) {
	reSymbol := regexp.MustCompile(`[^.1234567890]`)
	symbols := map[point]rune{}
	p1, p2 := 0, 0
	reNum := regexp.MustCompile(`\d+`)
	conns := map[point][]int{}

	for y, l := range s {
		for _, m := range reSymbol.FindAllStringIndex(l, -1) {
			symbols[point{y, m[0]}] = rune(s[y][m[0]])
		}
	}

	for y, l := range s {
		for _, m := range reNum.FindAllStringIndex(l, -1) {
		outter:
			for i := m[0]; i < m[1]; i++ {
				for _, next := range []point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
					p := point{y + next.y, i + next.x}
					if v, ok := symbols[p]; ok {
						intVal, _ := strconv.Atoi(l[m[0]:m[1]])
						fmt.Printf("%d connects to %c at %+v\n", intVal, v, p)
						p1 += intVal
						conns[p] = append(conns[p], intVal)
						break outter
					}
				}
			}
		}
	}

	for p, conn := range conns {
		if symbols[p] == '*' && len(conn) >= 2 {
			p2 += conn[1] * conn[0]
		}
	}

	fmt.Printf("P1:%d P2:%d\n", p1, p2)
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")
	schematic := []string{} // let's try treating them as strings
	schematic = append(schematic, inputLines...)

	sumParts2(schematic)
}
