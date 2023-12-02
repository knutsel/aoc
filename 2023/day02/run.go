package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseSubset(inp []string) (bool, int) {
	values := map[string]int{}
	maxValues := map[string]int{"red": 0, "green": 0, "blue": 0}
	isValid := true

	for _, s := range inp {
		f := strings.Split(s, ",")
		for i := 0; i < len(f); i++ {
			parts := strings.Fields(f[i])
			val, _ := strconv.Atoi(parts[0])
			maxValues[parts[1]] = max(maxValues[parts[1]], val)
			values[parts[1]] = val
		}

		if values["red"] > 12 || values["green"] > 13 || values["blue"] > 14 {
			isValid = false
		}
	}

	return isValid, maxValues["red"] * maxValues["green"] * maxValues["blue"]
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")
	sumP1, sumP2 := 0, 0

	for i, l := range inputLines {
		subsets := strings.Split(strings.Split(l, ":")[1], ";")

		ok, power := parseSubset(subsets)
		if ok {
			sumP1 += i + 1
		}

		sumP2 += power
	}

	fmt.Printf("P1:%d\nP2:%d\n", sumP1, sumP2)
}
