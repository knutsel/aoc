package day5

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseInitial(inpStr string) [][]rune {
	stacks := make([][]rune, 9)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")
	// for j := 1; j < 36; j += 4 {
	stackIndex := 0
	for j := 1; j < 36; j += 4 {
		fmt.Printf("stack:")
		stacks[stackIndex] = make([]rune, 0)
		for i := len(inputLines) - 2; i >= 0; i-- {
			fmt.Printf("[%d,%d]=%s ", i, j, string(inputLines[i][j]))
			if string(inputLines[i][j]) == " " {
				continue
			}
			stacks[stackIndex] = append(stacks[stackIndex], rune(inputLines[i][j]))
		}
		stackIndex++
		fmt.Println()
	}
	// 		fmt.Printf("in: %s\n", inputLines[i])
	// 		if i == 1 {
	// 			stacks[i] = make([]rune, 0)
	// 		}
	// 		stacks[i] = append(stacks[i], rune(inputLines[j][i]))
	// 	}
	// }

	return stacks
}

func printStacks(stacks [][]rune) {
	topRow := ""
	for _, row := range stacks {
		for _, c := range row {
			fmt.Printf(" %s", strconv.QuoteRune(c))
		}
		fmt.Println()
		if len(row) > 0 {
			topRow += string(row[len(row)-1])
		} else {
			topRow += " "
		}
	}
	fmt.Printf("top row %s\n", topRow)
}

func move(howmany, from, to int, stacks [][]rune) {
	fmt.Printf("Move  %d from %d to %d\n", howmany, from, to)
	// for l := 0; l < howmany; l++ {
	// 	stacks[to-1] = append(stacks[to-1], stacks[from-1][len(stacks[from-1])-1])
	// 	stacks[from-1] = stacks[from-1][:len(stacks[from-1])-1]
	// }
	stacks[to-1] = append(stacks[to-1], stacks[from-1][len(stacks[from-1])-howmany:]...)
	stacks[from-1] = stacks[from-1][:len(stacks[from-1])-howmany]
}

func Run() {
	fmt.Println("Starting Run (day 5)")

	inpBytes, _ := os.ReadFile("input.txt")
	inpStr := string(inpBytes)
	parts := strings.Split(strings.TrimSpace(inpStr), "\n\n")
	fmt.Printf("in: \n%s\n", parts[0])
	stacks := parseInitial(parts[0])
	fmt.Printf("s: %+v\n", stacks)
	printStacks(stacks)

	re := regexp.MustCompile("[0-9]+")
	for _, line := range strings.Split(parts[1], "\n") {
		nums := re.FindAllString(line, -1)
		howMany, _ := strconv.Atoi(nums[0])
		from, _ := strconv.Atoi(nums[1])
		to, _ := strconv.Atoi(nums[2])

		move(howMany, from, to, stacks)
		printStacks(stacks)
	}
	// fmt.Printf("contains=%d overlaps=%d\n", numContains, numOverlapping)
}
