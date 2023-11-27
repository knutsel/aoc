package day15b

import (
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"
)

const gridSize = int64(4_000_000)

type sensor struct {
	pos           point
	closestBeacon point
}

type linerange struct {
	start int64
	end   int64
}

type point struct {
	y int64
	x int64
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}

	return x
}

func mergeAndPruneRanges(inp []linerange) []linerange {
	// sort by start
	sort.Slice(inp, func(i, j int) bool {
		return inp[i].start < inp[j].start
	})

	merged := make([]linerange, 0)
	merged = append(merged, inp[0])

	for i := 1; i < len(inp); i++ {
		previousRange := merged[len(merged)-1]
		nextRange := inp[i]

		if nextRange.end < previousRange.end {
			continue // next fits in previous the starts are ordered
		}

		if nextRange.end >= previousRange.end && nextRange.start <= previousRange.end+1 {
			newRange := linerange{start: previousRange.start, end: nextRange.end}
			merged[len(merged)-1] = newRange // replace that slot with combined.

			continue
		}

		merged = append(merged, inp[i])
	}

	return merged
}

func Run() {
	fmt.Println("Starting Run (day 15)")

	inpBytes, _ := os.ReadFile("./input.txt")
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")

	// sList := make(map[point]sensor, 0)
	sList := make([]sensor, 0)

	for _, l := range inputLines {
		var sx, sy, bx, by int64

		// Sensor at x=20, y=1: closest beacon is at x=15, y=3
		_, err := fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d\n", &sx, &sy, &bx, &by)
		if err != nil {
			panic(err)
		}

		sList = append(sList, sensor{pos: point{sy, sx}, closestBeacon: point{by, bx}})
	}

	for yRef := int64(0); yRef < gridSize; yRef++ {
		noBeaconRanges := make([]linerange, 0)
		for _, s := range sList {
			sx := s.pos.x
			sy := s.pos.y
			bx := s.closestBeacon.x
			by := s.closestBeacon.y

			distance := Abs(sx-bx) + Abs(sy-by)
			// fmt.Printf("Sensor @: %d, %d beacon @ %d, %d -> manhattan distance:%d\n", sx, sy, bx, by, distance)

			extra := distance - Abs(yRef-sy)

			if by == yRef {
				noBeaconRanges = append(noBeaconRanges, linerange{start: bx, end: bx}) // the beacon itself
			}
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

				// fmt.Printf("Will cross: line %d with extra:%d at sx:%d center %+v\n", yRef, extra, sx, noBeaconRanges)
			}
		}

		m := mergeAndPruneRanges(noBeaconRanges)

		if len(m) == 1 && m[0].start <= 0 && m[0].end >= gridSize {
			// fmt.Printf(".")
		} else {
			fmt.Printf("\nline %d hasnobeacons %+v\n", yRef, m)
			fmt.Printf("answer is %d * %d +%d = %d\n", m[0].end, gridSize, yRef, m[0].end+1*gridSize+yRef)
			x := big.NewInt(m[0].end)
			y := big.NewInt(yRef)
			gs := big.NewInt(gridSize)
			mult := big.NewInt(0)
			ans := big.NewInt(0)

			mult.Mul(x, gs)

			fmt.Println("Answer:", ans.Add(mult, y))

			return
		}
	}
}
