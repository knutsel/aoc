package day10b

import (
	"fmt"
	"os"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Run() {
	fmt.Println("Starting Run (day 10)")
	inpBytes, _ := os.ReadFile("../day10/input.txt")
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")

	instrMap := make(map[string]int)
	instrMap["noop"] = 1
	instrMap["addx"] = 2

	X := 1
	cycle := 0
	var screen [6][40]string
	row := 0
	col := 0
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

			if Abs(X-col) <= 1 {
				screen[row][col] = "x"
			} else {
				screen[row][col] = "."
			}

			if instructionStr == "addx" && i == 1 {
				var add int
				fmt.Sscanf(operand, "%d", &add)
				X += add
			}

			col++
			if col%40 == 0 {
				col = 0
				row++
			}
		}
		fmt.Print("\n")
	}

	for _, row := range screen {
		for _, val := range row {
			fmt.Printf("%s", val)
		}
		fmt.Println()
	}
}
