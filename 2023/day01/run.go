package day01

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This is not as easy as it seems.... overlapping words like "twone" mean 2 as first, but 1 as last
// Golang doesn't support re's w positive lookahead, so hacking it
func translate(line string) string {
	trMap := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	toReplace := ""
	with := ""
	withLast := ""
	firstIndex := len(line)
	lastIndex := 0

	for k, v := range trMap {
		first := strings.Index(line, k)
		if first >= 0 && first < firstIndex {
			toReplace = k
			with = v
			firstIndex = first
		}
	}

	line = strings.Replace(line, toReplace, with, 1) // replace the first

	for k, v := range trMap {
		last := strings.LastIndex(line, k)
		if last > 0 && last > lastIndex {
			withLast = v
			lastIndex = last
		}
	}

	if lastIndex != 0 {
		line = line[:lastIndex] + withLast + line[lastIndex+len(withLast):] // replace the last, this doesn't actually remove the rest of the string
	}

	return line
}

func getSumOfDigits(inputLines []string, withSpelled bool) int {
	sum := 0

	for _, l := range inputLines {
		digits := []rune{}

		if withSpelled {
			l = translate(l)
		}

		for _, c := range l {
			if c >= '0' && c <= '9' {
				digits = append(digits, c)
			} else {
				continue
			}
		}

		num, err := strconv.Atoi(string(digits[0]))
		if err != nil {
			panic(err)
		}

		num1, err := strconv.Atoi(string(digits[len(digits)-1]))
		if err != nil {
			panic(err)
		}

		sum += 10*num + num1
	}

	return sum
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")

	fmt.Printf("P1:%d\n", getSumOfDigits(inputLines, false))
	fmt.Printf("P2:%d\n", getSumOfDigits(inputLines, true))
}
