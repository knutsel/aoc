package utils

import (
	"os"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	inpBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	inpStr := string(inpBytes)
	return strings.Split(strings.TrimSpace(inpStr), "\n"), nil
}
