package day09

import (
	"fmt"
	"os"
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

func calcSequences(sequence [][]int, level int) int {
	input := sequence[level]
	allZeros := true

	for _, v := range input {
		if v != 0 {
			allZeros = false
		}
	}

	if allZeros {
		for i := len(sequence) - 1; i > 0; i-- {
			sequence[i-1] = append(sequence[i-1], sequence[i][len(sequence[i])-1]+sequence[i-1][len(sequence[i-1])-1])
		}

		return sequence[0][len(sequence[0])-1]
	}

	diffSequence := []int{}
	for i := 0; i < len(input)-1; i++ {
		diffSequence = append(diffSequence, input[i+1]-input[i])
	}

	level++

	sequence = append(sequence, diffSequence)

	return calcSequences(sequence, level)
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
		p1 += calcSequences(s, 0)
	}

	fmt.Printf("P1: %d\n", p1)
	fmt.Printf("P2: %d\n", p2)
}
