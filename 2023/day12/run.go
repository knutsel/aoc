package day12

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rowInfo struct {
	springs string
	summary []int
}

func toInt(s string) int {
	iVal, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("%s:%s", s, err))
	}

	return iVal
}

// with a lot of help from https://github.com/ayoubzulfiqar/advent-of-code/blob/main/Go/Day12/part_1.go,
// started somewhat in the right direction, but got horribly lost
func numPossible(row string, summ []int) uint64 {
	if row == "" {
		if len(summ) == 0 {
			return 1
		}

		return 0
	}

	if len(summ) == 0 {
		if strings.Contains(row, "#") {
			return 0
		}

		return 1
	}

	count := uint64(0)

	if row[0] == '.' || row[0] == '?' {
		count += numPossible(row[1:], summ)
	}

	if row[0] == '#' || row[0] == '?' {
		if summ[0] <= len(row) && !strings.Contains(row[:summ[0]], ".") && (summ[0] == len(row) || row[summ[0]] != '#') {
			if summ[0] == len(row) {
				count += numPossible("", summ[1:])
			} else {
				count += numPossible(row[summ[0]+1:], summ[1:])
			}
		}
	}

	return count
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	p1, p2 := uint64(0), uint64(0)
	rows := []rowInfo{}

	for _, l := range strings.Split(strings.TrimSpace(string(inpBytes)), "\n") {
		summ := []int{}
		for _, s := range strings.Split(strings.Split(l, " ")[1], ",") {
			summ = append(summ, toInt(s))
		}

		ri := rowInfo{
			springs: strings.Split(l, " ")[0],
			summary: summ,
		}
		rows = append(rows, ri)
	}

	for _, r := range rows {
		p1 += numPossible(r.springs, r.summary)
		p2 += p1 * p1 * p1 * p1 * p1 * p1
	}

	fmt.Printf("P1: %d, P2:%d\n", p1, p2)
}
