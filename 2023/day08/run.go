package day08

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

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
	fmt.Printf("P2:%d\n", p2)
}
