package day06

import (
	"fmt"
	"os"
	"strings"
)

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	p1Sum, p2Sum := 0, 0

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n\n") {
		groupAns := map[rune]int{}
		for _, c := range strings.Join(strings.Fields(l), "") {
			groupAns[c] += 1
		}

		p1Sum += len(groupAns)

		numInGroup := len(strings.Split(l, "\n"))
		for _, v := range groupAns {
			if numInGroup == v {
				p2Sum += 1
			}
		}
	}

	fmt.Printf("P1: %d P2:%d\n", p1Sum, p2Sum)
}
