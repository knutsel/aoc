package day11b

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items           []int64
	activityCounter int
	op              struct {
		operation string
		arg1      int64
		arg2      int64
	}
	test struct {
		operation string
		arg1      int64
		yes       int
		no        int
	}
}

func doOp(oper string, arg1, arg2 int64) int64 {
	switch oper {
	case "+":
		return arg1 + arg2
	case "-":
		return arg1 - arg2
	case "*":
		return arg1 * arg2
	case "/":
		m := arg1 % arg2
		return m
	default:
		panic("Unknown op")
	}
}

func parseMonkeyInputStr(in string) monkey {
	m := monkey{
		items: make([]int64, 0),
	}

	mlines := strings.Split(in, "\n")
	for _, itemStr := range strings.Split(strings.Split(mlines[1], ":")[1], ",") {
		item, _ := strconv.ParseInt(strings.TrimSpace(itemStr), 10, 64)
		m.items = append(m.items, item)
	}
	optStr := strings.Fields(mlines[2])
	if optStr[0] != "Operation:" || optStr[1] != "new" {
		panic("NO op")
	}
	m.op.operation = optStr[4]
	m.op.arg1, _ = strconv.ParseInt(optStr[3], 10, 64)
	m.op.arg2, _ = strconv.ParseInt(optStr[5], 10, 64)

	testStr := strings.Fields(mlines[3])
	if testStr[0] != "Test:" {
		panic("NO TEST")
	}
	if testStr[1] == "divisible" {
		m.test.operation = "/"
		m.test.arg1, _ = strconv.ParseInt(testStr[3], 10, 64)
	}
	m.test.yes, _ = strconv.Atoi(strings.Fields(mlines[4])[5])
	m.test.no, _ = strconv.Atoi(strings.Fields(mlines[5])[5])

	return m
}

func getArgs(a1, a2, old int64) (int64, int64) {
	if a1 == 0 {
		a1 = old
	}
	if a2 == 0 {
		a2 = old
	}
	return a1, a2
}

func monkeyDo(monkeys []monkey, index int, worryDivisor int) {
	for i := range monkeys[index].items {
		monkeys[index].activityCounter++
		toInspect := monkeys[index].items[i]
		arg1, arg2 := getArgs(monkeys[index].op.arg1, monkeys[index].op.arg2, toInspect)
		newWorry := doOp(monkeys[index].op.operation, arg1, arg2)
		newWorry %= int64(worryDivisor)
		targetMonkey := monkeys[index].test.no
		if doOp(monkeys[index].test.operation, newWorry, monkeys[index].test.arg1) == 0 {
			targetMonkey = monkeys[index].test.yes
		}

		// fmt.Printf("tossing %d from %d to %d\n", newWorry, index, targetMonkey)
		monkeys[targetMonkey].items = append(monkeys[targetMonkey].items, newWorry)
	}
	monkeys[index].items = make([]int64, 0)
}

func Run() {
	fmt.Println("Starting Run (day 11)")
	inpBytes, _ := os.ReadFile("input.txt")
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n\n")

	monkeys := make([]monkey, 0)
	for _, monkeyStr := range inputLines {
		// fmt.Printf("%s", monkeyStr)
		m := parseMonkeyInputStr(monkeyStr)
		fmt.Printf("%+v\n", m)
		monkeys = append(monkeys, m)
	}

	// got this off the Internet... I need to think if I can make this make sense
	commonDivisor := 1
	for _, m := range monkeys {
		commonDivisor *= int(m.test.arg1)
	}
	for i := 0; i < 10000; i++ {
		for mNo := range monkeys {
			monkeyDo(monkeys, mNo, commonDivisor)
		}
	}

	act := make([]int, 0)
	for i, m := range monkeys {
		fmt.Printf("Monkey %d inspected %d items\n", i, m.activityCounter)
		act = append(act, m.activityCounter)
	}

	sort.Ints(act)

	fmt.Printf("%+v means: %d\n", act, act[(len(act)-2)]*act[(len(act)-1)])
}
