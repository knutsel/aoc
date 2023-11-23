package day13

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInt(in string) int {
	i, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}

	return i
}
func part1(depSchedStr []string, tStart int) {
	depSched := []int{}

	for _, v := range depSchedStr {
		if v == "x" {
			continue
		}

		depSched = append(depSched, toInt(v))
	}

	for t := tStart; ; t++ {
		for _, dt := range depSched {
			if t%dt == 0 {
				fmt.Printf("P1: %d\n", (t-tStart)*dt)
				return
			}
		}
	}
}

type schedEntry struct {
	startTime    int64
	diffFromPrev int64
}

func part2(depSchedStr []string, tStart int64) {
	depSched := []schedEntry{}
	fromPrev := 1

	for i, v := range depSchedStr {
		if v == "x" {
			fromPrev++
			continue
		}

		depSched = append(depSched, schedEntry{
			startTime:    int64(toInt(v)),
			diffFromPrev: int64(i),
		})

		fromPrev = 1
	}

	// got this from internet - i'm not that smart
	timeValue := int64(0)
	runningProduct := int64(1)

	for _, bus := range depSched {
		index, busID := bus.diffFromPrev, bus.startTime
		// this for loop adjusts the time until the constaint for this bus is met
		// i.e. ensure (time + index) is divisible by the busID to ensure the bus arrives
		for (timeValue+index)%busID != 0 {
			// running product is used to increment because it will not affect
			// the modulo of any of the previously scheduled busses, we've found
			// the frequency to match them.
			// e.g. if busID: 5 & index: 2, min timeValue is 3 b/c (3+2)%5 == 0
			//      if the running product were 5, adding 5 means (8+2)%5 == 0
			//      and (3 + 5x + 2) % 5 == 0 for any x
			timeValue += runningProduct
		}

		runningProduct *= busID
	}

	fmt.Printf("P2: %d\n", timeValue)
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	lines := strings.Split(strings.TrimSpace(inpStr), "\n")
	tStart := toInt(lines[0])
	depSchedStr := strings.Split(lines[1], ",")

	part1(depSchedStr, tStart)
	part2(depSchedStr, int64(tStart))
}
