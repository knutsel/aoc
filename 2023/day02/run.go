package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseSubset(inp []string) bool {
	r, g, b := 0, 0, 0

	for _, s := range inp {
		f := strings.Split(s, ",")
		for i := 0; i < len(f); i++ {
			parts := strings.Fields(f[i])

			val, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}

			switch parts[1] {
			case "red":
				r = val
			case "green":
				g = val
			case "blue":
				b = val
			}

			if r > 12 || g > 13 || b > 14 {
				return false
			}
		}
	}

	return true
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")
	// Game 68: 4 red, 2 blue, 5 green; 5 blue, 8 red, 2 green; 11 red, 2 green, 4 blue; 7 red, 5 blue, 3 green
	sum := 0

	for i, l := range inputLines {
		subsets := strings.Split(strings.Split(l, ":")[1], ";")
		if parseSubset(subsets) {
			// fmt.Printf("ID:%d is \n", i+1)
			sum += i + 1
		}
	}

	fmt.Printf("P1:%d\n", sum)
}
