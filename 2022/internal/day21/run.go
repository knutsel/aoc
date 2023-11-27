package day21

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// See day21b for a much cleaner recursive solution

func calc(in []string) int64 {
	a1, _ := strconv.ParseInt(in[0], 10, 64)
	a2, _ := strconv.ParseInt(in[2], 10, 64)

	switch in[1] {
	case "+":
		return a1 + a2
	case "-":
		return a1 - a2
	case "*":
		return a1 * a2
	case "/":
		return a1 / a2
	default:
		fmt.Printf("??? %q\n", in[2])
		return -1
	}
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := strings.TrimSpace(string(inpBytes))
	resolved := make(map[string][]string)

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		parts := strings.Fields(l)
		m := strings.TrimSuffix(parts[0], ":")
		resolved[m] = parts[1:]
	}

	for {
		if resolvedVal, ok := resolved["root"]; ok && len(resolvedVal) == 1 {
			fmt.Printf("root says %s\n", resolvedVal[0])
			break
		}

		for k, fields := range resolved {
			resCount := 0

			for i, field := range fields {
				if i%2 == 0 {
					_, err := strconv.Atoi(field)
					if err != nil { // NaN
						if resolvedVal, ok := resolved[field]; ok {
							if len(resolvedVal) == 1 {
								fields[i] = resolvedVal[0]
							}
						}
					} else { // aN
						resCount++
					}
				}
			}

			if resCount == 2 {
				newFields := make([]string, 1)
				newFields[0] = fmt.Sprintf("%d", calc(fields))
				fields = newFields
			}

			resolved[k] = fields
		}
	}
}
