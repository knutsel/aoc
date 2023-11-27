package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Run() {
	fmt.Println("Starting Run (day 1)")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	max := 0
	elve := 1
	allTotals := make([]int, 0)
	for scanner.Scan() {
		inp := scanner.Text()
		if inp != "" {
			num, err := strconv.Atoi(inp)
			if err != nil {
				log.Fatal(err)
			}
			sum += num
		} else {
			elve += 1
			allTotals = append(allTotals, sum)
			if max < sum {
				max = sum
			}
			sum = 0
		}
	}
	sort.Ints(allTotals)
	topThree := allTotals[len(allTotals)-3:]
	topThreeSum := 0
	for _, s := range topThree {
		topThreeSum += s
	}

	fmt.Printf("Max is %d there were %d elves, top 3 %+v, top 3 sum %d\n", max, elve, topThree, topThreeSum)
}
