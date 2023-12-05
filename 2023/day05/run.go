package day05

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type mapEntry struct {
	destination int
	source      int
	length      int
}

// global is faster
var mapOfEverything [][]mapEntry

func toInt(s string) int {
	iVal, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("%s:%s", s, err))
	}

	return iVal
}

func locForSeed(sourceVal int) int {
	for _, m := range mapOfEverything {
		for _, entry := range m {
			if sourceVal >= entry.source && sourceVal < entry.source+entry.length {
				sourceVal = entry.destination - entry.source + sourceVal // the value for the next mapping
				break
			}
		}
	}

	return sourceVal
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	seeds := []int{}
	mapOfEverything = [][]mapEntry{}

	for i, part := range strings.Split(strings.TrimSpace(inpStr), "\n\n") {
		if i == 0 {
			for _, strVal := range strings.Fields(part)[1:] {
				seeds = append(seeds, toInt(strVal))
			}

			continue
		} // else

		entries := []mapEntry{}

		for _, eLine := range strings.Split(strings.TrimSpace(strings.Split(part, ":")[1]), "\n") {
			entryParts := strings.Fields(eLine)
			entries = append(entries, mapEntry{
				destination: toInt(entryParts[0]),
				source:      toInt(entryParts[1]),
				length:      toInt(entryParts[2]),
			})
		}

		sort.Slice(entries, func(i, j int) bool { // sorting it saves 50% on total exec time
			return entries[i].source < entries[j].source
		})

		mapOfEverything = append(mapOfEverything, entries)
	}

	minLoc := math.MaxInt

	for _, seed := range seeds {
		loc := locForSeed(seed)
		minLoc = min(minLoc, loc)
	}

	fmt.Printf("P1:%d\n", minLoc)
	minLoc = math.MaxInt

	for i := 0; i < len(seeds)-2; i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			loc := locForSeed(j)
			minLoc = min(minLoc, loc)
		}
	}

	fmt.Printf("P2:%d\n", minLoc)
}
