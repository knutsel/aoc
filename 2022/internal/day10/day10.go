package day10

import (
	"fmt"
	"os"
	"strings"
)

// type instruction struct {
// 	name   string
// 	cycles int
// }

func Run() {
	fmt.Println("Starting Run (day 10)")
	inpBytes, _ := os.ReadFile("input.txt")
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")

	instrMap := make(map[string]int)
	instrMap["noop"] = 1
	instrMap["addx"] = 2

	X := 1
	cycle := 1
	sum := 0
	for _, line := range inputLines {
		parts := strings.Split(line, " ")
		instructionStr := parts[0]
		operand := ""
		if len(parts) > 1 {
			operand = strings.Join(parts[1:], " ")
		}

		fmt.Printf("Going to %q on %q:  ", instructionStr, operand)

		for i := 0; i < instrMap[instructionStr]; i++ {
			cycle++
			fmt.Printf(" | cycle %d for %s(%d) X:%d", cycle, instructionStr, i, X)

			if instructionStr == "addx" && i == 1 {
				var add int
				fmt.Sscanf(operand, "%d", &add)
				X += add
			}
			switch cycle {
			case 20, 60, 100, 140, 180, 220:
				fmt.Printf("\nX at %d: %d, signal: %d\n", cycle, X, cycle*X)
				sum += cycle * X
			}
		}

		fmt.Print("\n")

	}
	fmt.Printf("sum: %d\n", sum)
}
