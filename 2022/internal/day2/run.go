package day2

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	fmt.Println("Starting Run (day 2)")

	inpBytes, _ := os.ReadFile("input.txt")
	inpStr := string(inpBytes)

	part1, part2 := 0, 0
	for i, line := range strings.Split(inpStr, "\n") {
		switch line {
		case "A X":
			part1 += 4
			part2 += 3
		case "A Y":
			part1 += 8
			part2 += 4
		case "A Z":
			part1 += 3
			part2 += 8
		case "B X":
			part1 += 1
			part2 += 1
		case "B Y":
			part1 += 5
			part2 += 5
		case "B Z":
			part1 += 9
			part2 += 9
		case "C X":
			part1 += 7
			part2 += 2
		case "C Y":
			part1 += 2
			part2 += 6
		case "C Z":
			part1 += 6
			part2 += 7
		default:
			fmt.Printf("Not sure what to do with %q at line %d\n", line, i)
		}
	}

	fmt.Printf("score is part1:%d part2:%d\n", part1, part2)
}
