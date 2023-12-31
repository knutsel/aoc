package day16

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type intRange struct {
	start int
	stop  int
}

type rule struct {
	name   string
	ranges []intRange
}

type ticket []int

func abs(inp int) int {
	if inp < 0 {
		return -inp
	}

	return inp
}

func parseInputRules(input string) []rule {
	rules := []rule{}

	for _, l := range strings.Split(strings.TrimSpace(input), "\n") {
		start1, start2, stop1, stop2 := 0, 0, 0, 0
		name := strings.Split(l, ":")[0]
		p2 := strings.TrimSpace(strings.Split(l, ":")[1])
		fmt.Sscanf(p2, "%d%d or %d%d", &start1, &stop1, &start2, &stop2)
		r := rule{
			name:   name,
			ranges: []intRange{{start1, abs(stop1)}, {start2, abs(stop2)}},
		}
		rules = append(rules, r)
	}

	return rules
}

// [0] is my ticket
func parseTickets(input []string) []ticket {
	tickets := []ticket{}

	for _, tString := range input {
		for i, l := range strings.Split(strings.TrimSpace(tString), "\n") {
			if i == 0 {
				continue
			}

			ticket := []int{}

			for _, itemString := range strings.Split(l, ",") {
				item, _ := strconv.Atoi(itemString)
				ticket = append(ticket, item)
			}

			tickets = append(tickets, ticket)
		}
	}

	return tickets
}

func mergeRules(inp []rule) []intRange {
	ranges := []intRange{}
	for _, rule := range inp {
		ranges = append(ranges, rule.ranges...)
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	output := []intRange{}
	previous := ranges[0]

	for i := 1; i < len(ranges); i++ {
		next := ranges[i]
		if next.start <= previous.stop {
			if next.stop >= previous.stop {
				previous.stop = next.stop
			}
		} else {
			output = append(output, intRange{previous.start, previous.stop})
		}
	}

	output = append(output, intRange{previous.start, previous.stop})

	return output
}

// nolint: gocognit
func resolveTicketLayout(rules []rule, tickets []ticket) int {
	resolvedTickIndexes := map[int]bool{}
	resolvedRules := map[string]int{}

	for len(resolvedRules) < len(rules) {
		for loc := range tickets[0] {
			if resolvedTickIndexes[loc] {
				continue
			}

			valiedRuleNames := []string{}

			for _, r := range rules {
				if _, ok := resolvedRules[r.name]; ok {
					continue
				}

				allValid := true

				for _, t := range tickets {
					if !((t[loc] >= r.ranges[0].start && t[loc] <= r.ranges[0].stop) || (t[loc] >= r.ranges[1].start && t[loc] <= r.ranges[1].stop)) {
						allValid = false
						break
					}
				}

				if allValid {
					valiedRuleNames = append(valiedRuleNames, r.name)
				}
			}

			if len(valiedRuleNames) == 1 {
				resolvedTickIndexes[loc] = true
				resolvedRules[valiedRuleNames[0]] = loc
			}
		}
	}

	mult := 1

	for k, v := range resolvedRules {
		if strings.HasPrefix(k, "departure") {
			mult *= tickets[0][v]
		}
	}

	return mult
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	parts := strings.Split(strings.TrimSpace(inpStr), "\n\n")
	rules := parseInputRules(parts[0])
	tickets := parseTickets(parts[1:])
	validRanges := mergeRules(rules)
	sumVal := 0
	validTickets := []ticket{}

	for _, t := range tickets {
		valid := true
	tdone:
		for _, item := range t {
			for _, r := range validRanges {
				if item < r.start || item > r.stop {
					sumVal += item
					valid = false
					break tdone
				}
			}
		}

		if valid {
			validTickets = append(validTickets, t)
		}
	}

	fmt.Printf("P1:%d len(validTickets): %d\n", sumVal, len(validTickets))
	fmt.Printf("P2:%d\n", resolveTicketLayout(rules, validTickets))
}
