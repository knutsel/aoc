package day09

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func toInt(s string) int {
	iVal, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("%s:%s", s, err))
	}

	return iVal
}

func calcSequences(sequence [][]int, level int) (int, int) {
	if slices.Max(sequence[level]) == 0 && slices.Min(sequence[level]) == 0 {
		for i := len(sequence) - 1; i > 0; i-- {
			sequence[i-1] = append(sequence[i-1], sequence[i][len(sequence[i])-1]+sequence[i-1][len(sequence[i-1])-1])
			sequence[i-1] = append([]int{sequence[i-1][0] - sequence[i][0]}, sequence[i-1]...)
		}

		return sequence[0][len(sequence[0])-1], sequence[0][0]
	}

	diffSequence := []int{}
	for i := 0; i < len(sequence[level])-1; i++ {
		diffSequence = append(diffSequence, sequence[level][i+1]-sequence[level][i])
	}
	level++

	return calcSequences(append(sequence, diffSequence), level)
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	p1, p2 := 0, 0
	inp := [][][]int{}

	for _, l := range strings.Split(strings.TrimSpace(string(inpBytes)), "\n") {
		sequence := [][]int{}
		initialSequence := []int{}

		for _, strVal := range strings.Fields(l) {
			initialSequence = append(initialSequence, toInt(strVal))
		}

		sequence = append(sequence, initialSequence)
		inp = append(inp, sequence)
	}

	for _, s := range inp {
		p1val, p2Val := calcSequences(s, 0)
		p1 += p1val
		p2 += p2Val
	}

	fmt.Printf("P1: %d\nP2: %d\n", p1, p2)
}
