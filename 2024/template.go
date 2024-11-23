package utils

import (
	"fmt"
	"log"
)

func Template(fName string) {
	lines, err := ReadLines(fName)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		fmt.Printf("%s\n", line)
	}
}
