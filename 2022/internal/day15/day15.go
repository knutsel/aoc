package day15

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// const yRef = 10

const yRef = 2000000

type sensor struct {
	pos           point
	closestBeacon point
}

type linerange struct {
	start int
	end   int
}

type point struct {
	y int
	x int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func mergeRanges(inp []linerange) []linerange {
	// sort by start
	sort.Slice(inp, func(i, j int) bool {
		return inp[i].start < inp[j].start
	})

	merged := make([]linerange, 0)
	merged = append(merged, inp[0])

	for i := 1; i < len(inp); i++ {
		lastMerged := merged[len(merged)-1]
		if inp[i].start <= lastMerged.end {
			mRange := linerange{start: lastMerged.start, end: inp[i].end}
			merged[len(merged)-1] = mRange
		} else {
			merged = append(merged, inp[i])
		}
	}

	return merged
}

func Run() {
	fmt.Println("Starting Run (day 15)")

	inpBytes, _ := os.ReadFile("./input.txt")
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")

	noBeaconRanges := make([]linerange, 0)

	for _, l := range inputLines {
		var sx, sy, bx, by int

		// Sensor at x=20, y=1: closest beacon is at x=15, y=3
		_, err := fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d\n", &sx, &sy, &bx, &by)
		if err != nil {
			panic(err)
		}

		distance := Abs(sx-bx) + Abs(sy-by)
		fmt.Printf("Sensor @: %d, %d beacon @ %d, %d -> manhattan distance:%d\n", sx, sy, bx, by, distance)

		extra := distance - Abs(yRef-sy)

		if extra >= 0 {
			if yRef == sy && yRef == by {
				end := bx
				start := sx

				if start > end {
					end = sx
					start = bx
				}
				noBeaconRanges = append(noBeaconRanges, linerange{start: start, end: end})
			} else {
				noBeaconRanges = append(noBeaconRanges, linerange{start: sx - extra, end: sx + extra})
			}

			fmt.Printf("Will cross: %d at sx:%d center %+v\n", extra, sx, noBeaconRanges)
		}
	}

	m := mergeRanges(noBeaconRanges)
	sum := 0

	for _, r := range m {
		sum += r.end - r.start
	}

	fmt.Printf("line %d has %d nobeacons\n", yRef, sum)
}
