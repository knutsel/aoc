package day8

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type tree struct {
	height    int
	isVisible bool
}

const size = 99

func printForest(in [size][size]tree) {
	fmt.Print("\n")
	visNum := 0
	for _, row := range in {
		for _, val := range row {
			if val.isVisible {
				color.Set(color.FgRed)
				fmt.Printf("%d", val.height)
				color.Unset()
				visNum++
			} else {
				fmt.Printf("%d", val.height)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("visible:%d\n", visNum)
}

func markVisible(in [size][size]tree) {
	for i := 1; i < len(in)-1; i++ {
		highest := in[i][0].height
		for j := 1; j < len(in[i])-2; j++ {
			if in[i][j].height > highest {
				in[i][j].isVisible = true
				highest = in[i][j].height
			}
		}
	}
	for i := len(in) - 2; i > 0; i-- {
		highest := in[i][len(in[i])-1].height
		for j := len(in[i]) - 2; j > 0; j-- {
			if in[i][j].height > highest {
				in[i][j].isVisible = true
				highest = in[i][j].height
			}
		}
	}

	for j := 1; j < len(in)-2; j++ {
		highest := in[len(in)-1][j].height
		for i := len(in) - 2; i > 0; i-- {
			if in[i][j].height > highest {
				in[i][j].isVisible = true
				highest = in[i][j].height
			}
		}
	}

	for j := len(in) - 2; j > 0; j-- {
		highest := in[0][j].height
		for i := 1; i < len(in)-1; i++ {
			if in[i][j].height > highest {
				in[i][j].isVisible = true
				highest = in[i][j].height
			}
		}
	}

	printForest(in)
}

func Run(fName string) {
	fmt.Printf("Starting Run (day 8 w %s)\n", fName)

	forest := [size][size]tree{}
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")
	for i, line := range inputLines {
		for j, r := range line {
			vis := false
			if i == 0 || i == size-1 || j == 0 || j == size-1 {
				vis = true
			}
			forest[i][j] = tree{height: int(r - '0'), isVisible: vis}
		}
	}
	printForest(forest)

	markVisible(forest)
	// fmt.Printf("part1 :%d,\n", part1)
}
