package day06

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInt(s string) int {
	iVal, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("%s:%s", s, err))
	}

	return iVal
}

func race(t, d int) int {
	speed := 0
	won := 0
	timeLeft := t

	for timeLeft > 0 {
		speed++
		timeLeft--

		if timeLeft*speed > d {
			won++
		}
	}

	return won
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	raceInfo := [][]int{} // 0 is the times, 1 is the distances

	for _, l := range strings.Split(strings.TrimSpace(string(inpBytes)), "\n") {
		nums := []int{}
		for _, w := range strings.Fields(l)[1:] {
			nums = append(nums, toInt(w))
		}

		raceInfo = append(raceInfo, nums)
	}

	mult := 1

	for i := range raceInfo[0] {
		mult *= race(raceInfo[0][i], raceInfo[1][i])
	}

	fmt.Printf("P1:%d\n", mult)
}
