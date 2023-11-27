package day21b

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var resolved = map[string]string{}

func calRecursive(expr string) int {
	if v, err := strconv.Atoi(resolved[expr]); err == nil {
		return v
	}

	f := strings.Fields(resolved[expr])
	switch f[1] {
	case "+":
		return calRecursive(f[0]) + calRecursive(f[2])
	case "-":
		return calRecursive(f[0]) - calRecursive(f[2])
	case "*":
		return calRecursive(f[0]) * calRecursive(f[2])
	case "/":
		return calRecursive(f[0]) / calRecursive(f[2])
	}

	fmt.Printf("Shouldn't be here!\n")

	return -2
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := strings.TrimSpace(string(inpBytes))

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		parts := strings.Fields(l)
		m := strings.TrimSuffix(parts[0], ":")
		resolved[m] = strings.Join(parts[1:], " ")
	}

	part1 := calRecursive("root")

	// with a little help from the Internet; just iterating over all possible values for humn is too slow.
	// sort.Find() homes in on the answer remarkably fast
	resolved["humn"] = "0"
	s := strings.Fields(resolved["root"])

	if calRecursive(s[0]) < calRecursive(s[2]) {
		s[0], s[2] = s[2], s[0]
	}

	part2, _ := sort.Find(1e16, func(v int) int {
		resolved["humn"] = strconv.Itoa(v)
		fmt.Printf("Trying: %d\n", v)
		return calRecursive(s[0]) - calRecursive(s[2])
	})
	fmt.Printf("p1: %d p2: %d\n", part1, part2)
}
