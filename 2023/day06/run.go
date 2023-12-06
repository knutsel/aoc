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
	p1Info := [][]int{} // 0 is the times, 1 is the distances
	p2Info := []int{}
	p1 := 1

	for _, l := range strings.Split(strings.TrimSpace(string(inpBytes)), "\n") {
		nums := []int{}
		for _, w := range strings.Fields(l)[1:] {
			nums = append(nums, toInt(w))
		}

		p2Info = append(p2Info, toInt(strings.ReplaceAll(strings.Split(l, ":")[1], " ", "")))
		p1Info = append(p1Info, nums)
	}

	for i := range p1Info[0] {
		p1 *= race(p1Info[0][i], p1Info[1][i])
	}

	fmt.Printf("P1:%d P2:%d\n", p1, race(p2Info[0], p2Info[1]))
}
