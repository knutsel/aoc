package day11

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items           []int
	activityCounter int
	op              struct {
		operation string
		arg1      int
		arg2      int
	}
	test struct {
		operation string
		arg1      int
		yes       int
		no        int
	}
}

func doOp(oper string, arg1, arg2 int) int {
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
		items: make([]int, 0),
	}

	mlines := strings.Split(in, "\n")
	for _, itemStr := range strings.Split(strings.Split(mlines[1], ":")[1], ",") {
		item, _ := strconv.Atoi(strings.TrimSpace(itemStr))
		m.items = append(m.items, item)
	}
	optStr := strings.Fields(mlines[2])
	if optStr[0] != "Operation:" || optStr[1] != "new" {
		panic("NO op")
	}
	m.op.operation = optStr[4]
	m.op.arg1, _ = strconv.Atoi(optStr[3])
	m.op.arg2, _ = strconv.Atoi(optStr[5])

	testStr := strings.Fields(mlines[3])
	if testStr[0] != "Test:" {
		panic("NO TEST")
	}
	if testStr[1] == "divisible" {
		m.test.operation = "/"
		m.test.arg1, _ = strconv.Atoi(testStr[3])
	}
	m.test.yes, _ = strconv.Atoi(strings.Fields(mlines[4])[5])
	m.test.no, _ = strconv.Atoi(strings.Fields(mlines[5])[5])

	return m
}

func getArgs(a1, a2, old int) (int, int) {
	if a1 == 0 {
		a1 = old
	}
	if a2 == 0 {
		a2 = old
	}
	return a1, a2
}

func monkeyDo(monkeys []monkey, index int) {
	for i, _ := range monkeys[index].items {
		monkeys[index].activityCounter++
		toInspect := monkeys[index].items[i]
		arg1, arg2 := getArgs(monkeys[index].op.arg1, monkeys[index].op.arg2, toInspect)
		newWorry := doOp(monkeys[index].op.operation, arg1, arg2)
		newWorry = newWorry / 3
		targetMonkey := monkeys[index].test.no
		if doOp(monkeys[index].test.operation, newWorry, monkeys[index].test.arg1) == 0 {
			targetMonkey = monkeys[index].test.yes
		}

		fmt.Printf("tossing %d from %d to %d\n", newWorry, index, targetMonkey)
		monkeys[targetMonkey].items = append(monkeys[targetMonkey].items, newWorry)
	}
	monkeys[index].items = make([]int, 0)
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

	for i := 0; i < 20; i++ {
		for mNo := range monkeys {
			monkeyDo(monkeys, mNo)
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
