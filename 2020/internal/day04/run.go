package day04

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseInput(in string) map[string]string {
	out := make(map[string]string)

	for _, f := range strings.Fields(strings.ReplaceAll(in, "\n", " ")) {
		parts := strings.Split(f, ":")
		out[parts[0]] = parts[1]
	}

	return out
}

func check(b []map[string]string) {
	badCountP1 := 0
	badCountP2 := 0

	for _, passport := range b {
	b1:
		for _, f := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
			if _, ok := passport[f]; !ok {
				badCountP1++
				break b1
			} else if !checkP2(f, passport[f]) {
				badCountP2++
				break b1

			}
		}
	}

	fmt.Printf("Part1: %d Part2:%d\n", len(b)-badCountP1, len(b)-badCountP1-badCountP2)
}

func checkNum(value string, min, max int) bool {
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	if i >= min && i <= max {
		return true
	}

	return false
}

func checkRE(in string, re *regexp.Regexp) bool {
	return re.FindString(in) != ""
}

func checkP2(f string, value string) bool {
	switch f {
	case "byr":
		return checkNum(value, 1920, 2002)
	case "iyr":
		return checkNum(value, 2010, 2020)
	case "eyr":
		return checkNum(value, 2020, 2030)
	case "hgt":
		if strings.HasSuffix(value, "cm") {
			return checkNum(strings.TrimSuffix(value, "cm"), 150, 193)
		}

		if strings.HasSuffix(value, "in") {
			return checkNum(strings.TrimSuffix(value, "in"), 59, 76)
		}

		return false
	case "hcl":
		return checkRE(value, regexp.MustCompile("^#[0-9a-f]{1,6}$"))
	case "ecl":
		return checkRE(value, regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$"))
	case "pid":
		return checkRE(value, regexp.MustCompile(`^\d{9}$`))
	}

	return false
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	batch := []map[string]string{}

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n\n") {
		passport := parseInput(l)
		batch = append(batch, passport)
	}

	check(batch)
}
