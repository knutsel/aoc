package day6

import (
	"fmt"
	"os"
)

func Run() {
	fmt.Println("Starting Run (day 6)")

	inpBytes, _ := os.ReadFile("input.txt")
	inpStr := string(inpBytes)

	for i := 0; i < len(inpStr)-14; i++ {
		seen := make(map[string]bool)
		for j := i; j < i+14; j++ {
			seen[string(inpStr[j])] = true
		}
		// fmt.Printf("i:%d len:%d\n", i, len(seen))
		if len(seen) == 14 {
			fmt.Printf("14 at %d\n", i+14)
			break
		}
	}
}
