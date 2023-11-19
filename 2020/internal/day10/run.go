package day10

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var cache = map[int]int{}

func findArrangements(start int, inp []int) int {
	if v, ok := cache[start]; ok {
		return v
	}

	if start == len(inp)-2 {
		return 1
	}

	sum := 0

	for i := start + 1; i < len(inp)-1; i++ {
		if inp[i]-inp[start] <= 3 {
			sum += findArrangements(i, inp)
		} else {
			break
		}
	}

	cache[start] = sum

	return sum
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	inp := []int{}

	inp = append(inp, 0) // the wall

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		intVal := 0
		fmt.Sscanf(l, "%d", &intVal)
		inp = append(inp, intVal)
	}

	sort.Ints(inp)

	inp = append(inp, inp[len(inp)-1]+3) // my device
	jumps := map[int]int{}

	for i := 0; i < len(inp)-1; i++ {
		diff := inp[i+1] - inp[i]
		jumps[diff]++
	}

	fmt.Printf("P1:%d\nP2:%d\n", jumps[1]*jumps[3], findArrangements(0, inp))
}
