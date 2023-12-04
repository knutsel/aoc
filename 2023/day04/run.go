package day04

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	y int
	x int
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	sum, sumP2 := 0, 0
	copies := map[int]int{}

	for i, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		parts := strings.Split(l, "|")
		winningNums := strings.Fields(strings.Split(parts[0], ":")[1])
		myNums := strings.Fields(parts[1])
		winning := map[string]bool{}

		for _, n := range winningNums {
			winning[n] = true
		}

		thisCard := 0

		for _, val := range myNums {
			if winning[val] {
				thisCard++
			}
		}

		if thisCard > 0 {
			sum += 1 << (thisCard - 1)
		}

		copies[i]++
		sumP2++

		// had a little help with this.
		for j := 1; j <= thisCard; j++ {
			copies[i+j] += copies[i]
			sumP2 += copies[i]
		}
	}

	fmt.Printf("P1:%d P2:%d\n", sum, sumP2)
}
