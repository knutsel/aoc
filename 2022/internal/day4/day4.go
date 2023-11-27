package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type secList struct {
	start int
	end   int
}

func getList(inp string) secList {
	part := strings.Split(inp, "-")
	start, _ := strconv.Atoi(part[0])
	stop, _ := strconv.Atoi(part[1])
	return secList{
		start: start,
		end:   stop,
	}
}

func contains(first, second secList) bool {
	if first.start >= second.start && first.end <= second.end {
		return true
	}
	if second.start >= first.start && second.end <= first.end {
		return true
	}
	return false
}

func overlaps(first, second secList) bool {
	if second.start >= first.start && second.start <= first.end {
		return true
	}
	if first.start >= second.start && first.start <= second.end {
		return true
	}
	return false
}

func Run() {
	fmt.Println("Starting Run (day 4)")

	inpBytes, _ := os.ReadFile("input.txt")
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")

	numOverlapping := 0
	numContains := 0

	for _, line := range inputLines {
		ranges := strings.Split(line, ",")
		range1 := getList(ranges[0])
		range2 := getList(ranges[1])
		if contains(range1, range2) {
			numContains += 1
		}
		if overlaps(range1, range2) {
			numOverlapping += 1
		}

	}
	fmt.Printf("contains=%d overlaps=%d\n", numContains, numOverlapping)
}
