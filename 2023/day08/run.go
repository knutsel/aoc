package day08

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	p1, p2 := 0, 0
	right := map[string]string{}
	left := map[string]string{}
	re := regexp.MustCompile(`^(?P<key>\w{3}) = \((?P<left>\w{3}), (?P<right>\w{3})\)$`)
	leftRight := strings.Split(strings.TrimSpace(string(inpBytes)), "\n")[0]

	for _, l := range strings.Split(strings.TrimSpace(string(inpBytes)), "\n")[2:] {
		mapping := re.FindStringSubmatch(l)
		left[mapping[1]] = mapping[2]
		right[mapping[1]] = mapping[3]
	}

	current := "AAA"

	for {
		side := leftRight[p1%len(leftRight)]
		p1++

		if side == 'L' {
			current = left[current]
		} else {
			current = right[current]
		}

		if current == "ZZZ" {
			break
		}
	}

	fmt.Printf("P1:%d\n", p1)

	current = "A"
	currentList := []string{}
	for k, _ := range left {
		if strings.HasSuffix(k, current) {
			currentList = append(currentList, k)
		}
	}

	zCount := 0
	jumps := []int{}
outter:
	for {
		side := leftRight[p2%len(leftRight)]
		p2++
		for i, _ := range currentList {
			if side == 'L' {
				currentList[i] = left[currentList[i]]
			} else {
				currentList[i] = right[currentList[i]]
			}

			if strings.HasSuffix(currentList[i], "Z") {
				zCount++
				fmt.Printf("Z suffix at %d for index %d zCount:%d (len:%d)\n", p2, i, zCount, len(currentList))
				jumps = append(jumps, p2)
			}
			if zCount == len(currentList) {
				break outter
			}
		}
	}

	fmt.Printf("P2:%d\n", LCM(jumps[0], jumps[1], jumps[2:]...))
}
