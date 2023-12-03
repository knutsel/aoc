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

func sumParts(s []string) int {
	reSymbol := regexp.MustCompile(`[^.1234567890]`)
	isNeighbor := map[point]bool{}

	for y, l := range s {
		for _, x := range reSymbol.FindAllStringIndex(l, -1) {
			for _, next := range []point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
				p := point{y + next.y, x[0] + next.x}
				isNeighbor[p] = true
			}
		}
	}

	sum := 0
	reNum := regexp.MustCompile(`\d+`)

	for y, l := range s {
		for _, m := range reNum.FindAllStringIndex(l, -1) {
			if isNeighbor[point{y, m[0]}] || isNeighbor[point{y, m[1] - 1}] {
				numVal, err := strconv.Atoi(l[m[0]:m[1]])
				if err != nil {
					panic(err)
				}
				// fmt.Printf("Adding %d\n", numVal)
				sum += numVal
			}
		}
	}

	return sum
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")
	schematic := []string{} // let's try treating them as strings
	schematic = append(schematic, inputLines...)

	fmt.Printf("P1:%d\n", sumParts(schematic))
}
