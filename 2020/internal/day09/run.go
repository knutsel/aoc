package day09

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func checkXMAS(inp []int, plen int) int {
	for i := plen; i < len(inp); i++ {
		valid := false
	out:
		for ia := i - plen; ia < i; ia++ {
			for ib := i - plen; ib < i; ib++ {
				if inp[i] == inp[ia]+inp[ib] {
					valid = true
					break out
				}
			}
		}

		if !valid {
			return i
		}
	}

	return -1
}

func checkXMASP2(inp []int, startIndex int) int {
	for i := startIndex - 1; i > 0; i-- {
		sum := 0
		for j := i; j > 0; j-- {
			sum += inp[j]
			if sum > inp[startIndex] {
				break
			}

			if sum == inp[startIndex] {
				sort.Ints(inp[j : i+1])
				return inp[j] + inp[i]
			}
		}
	}

	return -1
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	inp := []int{}

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		intVal := 0
		fmt.Sscanf(l, "%d", &intVal)
		inp = append(inp, intVal)
	}

	indexP1 := checkXMAS(inp, 25)

	fmt.Printf("P1: %d\n", inp[indexP1])
	fmt.Printf("P2: %d\n", checkXMASP2(inp, indexP1))
}
