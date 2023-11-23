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

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	lines := strings.Split(strings.TrimSpace(inpStr), "\n")
	tStart := toInt(lines[0])
	depSchedStr := strings.Split(lines[1], ",")
	depSched := []int{}

	for _, v := range depSchedStr {
		if v == "x" {
			continue
		}

		depSched = append(depSched, toInt(v))
	}

	for t := tStart; ; {
		for _, dt := range depSched {
			if t%dt == 0 {
				fmt.Printf("break at t=%d, dt=%d --> P1:%d \n", t, dt, (t-tStart)*dt)
				return
			}
		}
		t++
	}
}
