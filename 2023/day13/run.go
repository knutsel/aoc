package day13

import (
	"fmt"
	"os"
	"strings"
)

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	p1 := 0
	for _, g := range strings.Split(strings.TrimSpace(string(inpBytes)), "\n\n") {
		p1 += summarize(g)
	}
	fmt.Printf("P1: %d, P2:%d\n", p1, p1)
}
