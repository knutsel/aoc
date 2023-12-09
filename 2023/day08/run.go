package day08

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// GCD, LCM from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int { // greatest common divisor (GCD) via Euclidean algorithm
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func LCM(a, b int, integers ...int) int { // find Least Common Multiple (LCM) via GCD
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func follow(from string, to string, left, right map[string]string, leftRight string) int {
	currentList := []string{}
	for k, _ := range left {
		if strings.HasSuffix(k, from) {
			currentList = append(currentList, k)
		}
	}

	zCount, sum := 0, 0
	jumps := []int{}

outter:
	for {
		side := leftRight[sum%len(leftRight)]
		sum++
		for i := range currentList {
			if side == 'L' {
				currentList[i] = left[currentList[i]]
			} else {
				currentList[i] = right[currentList[i]]
			}

			if strings.HasSuffix(currentList[i], to) {
				zCount++
				jumps = append(jumps, sum)
			}
			if zCount == len(currentList) {
				break outter
			}
		}
	}

	if len(jumps) > 1 {
		return LCM(jumps[0], jumps[1], jumps[2:]...)
	}

	return jumps[0]
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	right := map[string]string{}
	left := map[string]string{}
	re := regexp.MustCompile(`^(?P<key>\w{3}) = \((?P<left>\w{3}), (?P<right>\w{3})\)$`)
	leftRight := strings.Split(strings.TrimSpace(string(inpBytes)), "\n")[0]

	for _, l := range strings.Split(strings.TrimSpace(string(inpBytes)), "\n")[2:] {
		mapping := re.FindStringSubmatch(l)
		left[mapping[1]] = mapping[2]
		right[mapping[1]] = mapping[3]
	}

	fmt.Printf("P1: %d\n", follow("AAA", "ZZZ", left, right, leftRight))
	fmt.Printf("P2: %d\n", follow("A", "Z", left, right, leftRight))
}
