package day7

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

func Run() {
	fmt.Println("Starting Run (day 7)")

	inpBytes, _ := os.ReadFile("input.txt")
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")

	cwd := "/"
	sizeByDir := make(map[string]int)
	for _, line := range inputLines {
		if strings.HasPrefix(line, "$ cd") {
			to := strings.Fields(line)[2]
			if to == ".." {
				cwd = filepath.Clean(filepath.Dir(cwd))
			} else {
				cwd = filepath.Clean(cwd + "/" + to)
				sizeByDir[cwd] = 0
			}
		}
		if unicode.IsDigit(rune(line[0])) {
			sz, _ := strconv.Atoi(strings.Fields(line)[0])
			sizeByDir[cwd] += sz
		}
	}
	totalSizeByDir := make(map[string]int)
	for path, size := range sizeByDir {
		parent := filepath.Dir(path)
		totalSizeByDir[path] += size
		if path == "/" {
			continue
		}
		for {
			totalSizeByDir[parent] += size
			if parent == "/" {
				break
			}
			parent = filepath.Dir(parent)

		}
	}
	sum := 0
	part2 := totalSizeByDir["/"]
	for _, sz := range totalSizeByDir {
		if sz <= 100000 {
			sum += sz
		}
		if sz+70000000-totalSizeByDir["/"] >= 30000000 && sz < part2 {
			part2 = sz
		}
	}
	fmt.Printf("part1 :%d, part2:%d\n", sum, part2)
}
