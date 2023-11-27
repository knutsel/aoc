package day20

import (
	"fmt"
	"os"
	"strings"
)

// Hmm.
type entry struct {
	val      int
	hasMoved bool
}

func move(in []entry) []entry {
	from := 0

	for i := 0; i < len(in); i++ {
		if !in[i].hasMoved {
			from = i
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

	to := (from + in[from].val) % (len(in) - 1)
	if to <= 0 {
		to = len(in) - 1 + to // It's negative
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

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		var n int

		fmt.Sscanf(l, "%d", &n)
		e := entry{
			val:      n,
			hasMoved: false,
		}
		input = append(input, e)
		output = append(output, e)
	}

	for i := range input {
		output = move(output)
		fmt.Printf("i:%d ", i)
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
