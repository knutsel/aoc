package day05

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type mapEntry struct {
	destination int
	source      int
	length      int
}

func toInt(s string) int {
	iVal, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("%s:%s", s, err))
	}

	return iVal
}

func locForSeed(mapOfEverything map[string][]mapEntry, sourceVal int) int {
	sourceName := "seed"

	for {
		for mName, m := range mapOfEverything {
			mFrom := strings.Split(mName, "-")[0]
			mTo := strings.Fields(strings.Split(mName, "-")[2])[0]

			if mFrom != sourceName {
				continue
			}

			sourceName = mTo // the value for the next mapping

			for _, entry := range m {
				if sourceVal >= entry.source && sourceVal <= entry.source+entry.length {
					sourceVal = entry.destination - entry.source + sourceVal // the value for the next mapping
					break
				}
			}

			if sourceName == "location" {
				return sourceVal
			}

			break
		}
	}
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	seeds := []int{}
	mapOfEverything := map[string][]mapEntry{}

	for i, part := range strings.Split(strings.TrimSpace(inpStr), "\n\n") {
		if i == 0 {
			for _, strVal := range strings.Fields(part)[1:] {
				seeds = append(seeds, toInt(strVal))
			}

			continue
		} // else

		name := strings.Split(part, ":")[0]
		entries := []mapEntry{}

		for _, eLine := range strings.Split(strings.TrimSpace(strings.Split(part, ":")[1]), "\n") {
			entryParts := strings.Fields(eLine)
			entries = append(entries, mapEntry{
				destination: toInt(entryParts[0]),
				source:      toInt(entryParts[1]),
				length:      toInt(entryParts[2]),
			})
		}

		// sort.Slice(entries, func(i, j int) bool {
		// 	return entries[i].source < entries[j].source
		// })

		mapOfEverything[name] = entries
	}

	minLoc := math.MaxInt

	for _, seed := range seeds {
		loc := locForSeed(mapOfEverything, seed)
		minLoc = min(minLoc, loc)
	}

	fmt.Printf("P1:%d\n", minLoc)
}
