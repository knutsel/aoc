package aoc2024
//

import (
	"os"
	"strings"
)

func readLines(path string) ([]string, error) {
	inpBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	inpStr := string(inpBytes)
	return strings.Split(strings.TrimSpace(inpStr), "\n"), nil
}
