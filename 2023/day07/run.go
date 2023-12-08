package day07

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type play struct {
	hand string
	bid  int
}

// nolint: gochecknoglobals
var order = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'} // add a 1 for dummy and to be able to * the index

func scoreHand(h string) uint64 {
	sum := uint64(0)

	for _, card := range order {
		count := strings.Count(h, string(card))
		switch count {
		case 3, 4, 5:
			sum += uint64(count) * 0x2000000 // 2 to make sure 3 of a kind > 2 pair
		case 2:
			sum += 0x1000000
		}
	}

	for i, handCard := range h {
		for j, card := range order {
			if handCard == card {
				if (4 - i) > 0 {
					sum += uint64(j) << ((4 - i) * 4)
				} else {
					sum += uint64(j)
				}
			}
		}
	}

	return sum
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	p1 := 0
	entries := []play{}

	for _, l := range strings.Split(strings.TrimSpace(string(inpBytes)), "\n") {
		p := play{}
		fmt.Sscanf(l, "%s %d", &p.hand, &p.bid)

		entries = append(entries, p)
	}

	sort.Slice(entries, func(i, j int) bool {
		return scoreHand(entries[i].hand) < scoreHand(entries[j].hand)
	})

	for i, hand := range entries {
		p1 += (i + 1) * hand.bid
	}

	fmt.Printf("P1:%d P2:%d\n", p1, p1)
}
