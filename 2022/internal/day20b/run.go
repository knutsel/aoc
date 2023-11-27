package day20b

import (
	"fmt"
	"os"
	"strings"
)

type entry struct {
	val           int64
	originalIndex int
	hasMoved      bool
}

const (
	multFactor = 811589153 // day2
	repeat     = 10        // day2
	// multFactor = 1 // day1
	// repeat     = 1  // day1
)

func move(index int, in []entry) []entry {
	from := int64(0)

	for i := 0; i < len(in); i++ {
		if in[i].originalIndex == index {
			from = int64(i)
			break
		}
	}

	in[from].hasMoved = true
	if in[from].val == 0 {
		return in
	}

	out := make([]entry, 0)
	out = append(out, in[:from]...)
	out = append(out, in[from+1:]...)

	to := (from + in[from].val) % (int64(len(in)) - 1)
	if to <= 0 {
		to = int64(len(in)) - 1 + to // It's negative
	}

	out1 := make([]entry, 0)
	out1 = append(out1, out[:to]...)
	out1 = append(out1, in[from])
	out1 = append(out1, out[to:]...)

	return out1
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := strings.TrimSpace(string(inpBytes))
	input := make([]entry, 0)
	output := make([]entry, 0)

	for i, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		var n int64

		fmt.Sscanf(l, "%d", &n)
		e := entry{val: n * multFactor, originalIndex: i, hasMoved: false}
		input = append(input, e)
		output = append(output, e)
	}

	for j := 0; j < repeat; j++ {
		for i := 0; i < len(input); i++ {
			output = move(i, output)
		}
	}

	zeroAt := 0

	for i := range output {
		if output[i].val == 0 {
			zeroAt = i
			break
		}
	}

	fmt.Printf("zeroAt:%d -> %d\n", zeroAt, output[zeroAt+1000].val+output[zeroAt+2000].val+output[zeroAt+3000].val)
}
