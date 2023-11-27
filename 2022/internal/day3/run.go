package day3

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func runeVal(inp rune) int {
	val := int(inp) - int(rune('a')) + 1

	if unicode.IsUpper(inp) {
		val = int(inp) - int(rune('A')) + 27
	}
	return val
}

func intersect(first, second string) string {
	intersection := ""

	seen := make(map[rune]bool)
	for _, c := range first {
		if seen[c] {
			continue
		}
		seen[c] = true
		if strings.ContainsRune(second, c) {
			intersection += string(c)
		}
	}

	return intersection
}

func Run(fName string) {
	fmt.Printf("Starting Run (day 3 (%s)\n", fName)

	p1Sum, p2Sum := 0, 0
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")

	for i, line := range inputLines {
		inBothSides := intersect(line[len(line)/2:], line[:len(line)/2])
		p1Sum += runeVal(rune(inBothSides[0]))

		if i%3 == 2 {
			common := intersect(inputLines[i], inputLines[i-1])
			common = intersect(common, inputLines[i-2])
			p2Sum += runeVal(rune(common[0]))
		}
	}
	fmt.Printf("p1Sum=%d, p2Sum=%d\n", p1Sum, p2Sum)
}
