package day01

import (
	"fmt"
	"os"
	"strings"
)

func Run(fName string) {
	inpBytes, _ := os.ReadFile("input.txt")
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")

	inputInts := []int{}

	for _, line := range inputLines {
		var intVal int
		fmt.Sscanf(line, "%d", &intVal)
		inputInts = append(inputInts, intVal)
	}
outter:
	for i := range inputInts {
		for j := range inputInts {
			if inputInts[i]+inputInts[j] == 2020 {
				fmt.Printf("Break at (%d,%d), with %d and %d ==> Part 1 answer %d\n", i, j, inputInts[i], inputInts[j], inputInts[i]*inputInts[j])
				break outter
			}
		}
	}
outter1:
	for i := range inputInts {
		for j := range inputInts {
			for k := range inputInts {
				if inputInts[i]+inputInts[j]+inputInts[k] == 2020 {
					fmt.Printf("Break at (%d,%d,%d), with %d, %d, %d ==> Part 2 answer %d\n", i, j, k, inputInts[i], inputInts[j], inputInts[k], inputInts[i]*inputInts[j]*inputInts[k])
					break outter1
				}
			}
		}
	}
}
