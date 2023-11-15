package day08

import (
	"fmt"
	"os"
	"strings"
)

type instruction struct {
	operation string
	argument  int
}

func runProgram(program []instruction) (bool, int) {
	accVal, pc := 0, 0
	visited := map[int]bool{}

	for {
		if ok := visited[pc]; ok {
			return true, accVal
		}

		visited[pc] = true

		switch program[pc].operation {
		case "acc":
			accVal += program[pc].argument
			pc++
		case "nop":
			pc++
		case "jmp":
			pc += program[pc].argument
		}

		if pc == len(program)-1 {
			return false, accVal
		}
	}
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	program := []instruction{}

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		strVal, intVal := "", 0
		fmt.Sscanf(l, "%s %d", &strVal, &intVal)
		program = append(program, instruction{operation: strVal, argument: intVal})
	}

	_, accVal := runProgram(program)
	fmt.Printf("P1:%d\n", accVal)

	// should have a test for more runs than jmps, also need 2 loops.
	patched := 0

	for {
		patchedProgram := []instruction{}
		nopNo := 0

		for i := range program {
			patchedProgram = append(patchedProgram, program[i])
			if patchedProgram[i].operation == "jmp" {
				if nopNo == patched {
					patchedProgram[i].operation = "nop"
				}
				nopNo++
			}
		}
		patched++

		loops, accVal := runProgram(patchedProgram)
		if !loops {
			fmt.Printf("P2:%d\n", accVal)
			break
		}
	}
}
