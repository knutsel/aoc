package day02

import (
	"fmt"
	"os"
	"strings"
)

func Run(fName string) {
	inpBytes, _ := os.ReadFile("input.txt")
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")

	part1Valid, part2Valid := 0, 0

	for _, line := range inputLines {
		fmt.Printf("%s ", line)
		var first, last int

		var c rune
		var p string
		fmt.Sscanf(line, "%d-%d %c: %s", &first, &last, &c, &p)
		o := strings.Count(p, string(c))
		if o >= first && o <= last {
			part1Valid++
		}
		if (p[first-1] == byte(c) || p[last-1] == byte(c)) && p[first-1] != p[last-1] {
			fmt.Printf(" 2V\n")
			part2Valid++
		}
	}
	fmt.Printf("part1:%d part2:%d\n", part1Valid, part2Valid)
}
