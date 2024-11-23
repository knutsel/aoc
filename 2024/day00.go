package aoc2024

import (
	"fmt"
	"log"
)

func Day00(fName string) {
	lines, err := readLines(fName)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		fmt.Printf("%s\n", line)
	}
}
