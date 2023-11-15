package day05

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func BSP2ID(in string) int {
	var id int

	in = strings.ReplaceAll(in, "F", "0")
	in = strings.ReplaceAll(in, "B", "1")
	in = strings.ReplaceAll(in, "R", "1")
	in = strings.ReplaceAll(in, "L", "0")

	_, err := fmt.Sscanf(in, "%b", &id)
	if err != nil {
		panic(err)
	}

	return id
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	maxID := 0
	all := []int{}

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		id := BSP2ID(l)
		maxID = max(maxID, id)
		all = append(all, id)
	}

	fmt.Printf("P1: %d\n", maxID)

	sort.Ints(all)

	for i, v := range all {
		if all[i+1] != v+1 {
			fmt.Printf("p2: (%d): %d\n", i, v+1)
			break
		}
	}
}
