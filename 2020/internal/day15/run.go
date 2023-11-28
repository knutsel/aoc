package day15

import "fmt"

func recite(rounds int) int {
	spoken := map[int][]int{}
	input := []int{14, 8, 16, 0, 1, 17}

	for i := 0; i < len(input); i++ {
		spoken[input[i]] = append(spoken[input[i]], i)
	}

	lastSpoken := input[len(input)-1]
	for i := len(input); i < rounds; i++ {
		if spokenAtArray, ok := spoken[lastSpoken]; ok {
			if len(spokenAtArray) == 1 {
				lastSpoken = 0

				spoken[0] = append(spoken[0], i)
			} else {
				lastSpoken = spokenAtArray[len(spokenAtArray)-1] - spokenAtArray[len(spokenAtArray)-2]
				spoken[lastSpoken] = append(spoken[lastSpoken], i)
			}
		}
	}

	return lastSpoken
}

func Run() {
	fmt.Printf("P1:%d, P2:%d\n", recite(220), recite(30000000))
}
