package day17b

import (
	"fmt"
	"os"
	"strings"
)

// this is terribly hacky :-)
// couldn't get it to work, stole off the Internet

func createRockList() [][]string {
	list := make([][]string, 0)
	list = append(list, []string{"..@@@@."})
	list = append(list, []string{"...@...", "..@@@..", "...@..."})
	list = append(list, []string{"....@..", "....@..", "..@@@.."})
	list = append(list, []string{"..@....", "..@....", "..@....", "..@...."})
	list = append(list, []string{"..@@...", "..@@..."})

	return list
}

type chamber struct {
	tunnel      []string
	rocks       [][]string
	blows       []rune
	currentBlow int
}

func (ch chamber) print() {
	for i := len(ch.tunnel) - 1; i >= 0; i-- {
		fmt.Printf("%s\n", ch.tunnel[i])
	}
}

func (ch *chamber) solidify() {
	for i := len(ch.tunnel) - 1; i >= 0; i-- {
		if strings.Contains(ch.tunnel[i], "@") {
			ch.tunnel[i] = strings.ReplaceAll(ch.tunnel[i], "@", "#")
		}
	}
}

func (ch *chamber) drop() bool {
	for i := len(ch.tunnel) - 1; i >= 0; i-- {
		if strings.Contains(ch.tunnel[i], "@") {
			if strings.Contains(ch.tunnel[i-1], "-") { // first block at the bottom
				ch.solidify()
				return true
			}

			if ch.tunnel[i-1] == "......." { // just empty row below rock
				ch.tunnel = append(ch.tunnel[:i-1], ch.tunnel[i:]...)
				return false
			}

			for j := range ch.tunnel[i] {
				if ch.tunnel[i][j] == '@' && ch.tunnel[i-1][j] == '#' {
					ch.solidify()
					return true
				}
			}
		}
	}

	for i := 1; i < len(ch.tunnel)-1; i++ {
		newRow := []rune(strings.ReplaceAll(ch.tunnel[i], "@", "."))

		for j := range ch.tunnel[i] {
			if ch.tunnel[i+1][j] == '@' {
				newRow[j] = '@'
			}
		}

		ch.tunnel[i] = string(newRow)
	}

	ch.tunnel[len(ch.tunnel)-1] = "......."

	return false
}

func moveLine(in string, dir string) string {
	inRunes := []rune(in)
	out := []rune(strings.ReplaceAll(in, "@", "."))

	switch dir {
	case "left":
		for i := 0; i < len(inRunes)-1; i++ {
			if inRunes[i+1] == '@' {
				out[i] = '@'
			}
		}
	case "right":
		for i := len(inRunes) - 1; i >= 0; i-- {
			if inRunes[i] == '@' {
				out[i+1] = '@'
			}
		}
	default:
		panic("invalid direction")
	}

	return string(out)
}

// move only gets called when not at the wall
func (ch *chamber) move(dir string) bool {
	for i := len(ch.tunnel) - 1; i >= 0; i-- {
		if strings.Contains(ch.tunnel[i], "@") {
			ch.tunnel[i] = moveLine(ch.tunnel[i], dir)
		}
	}

	return false
}

func (ch *chamber) dropNewRock(rockType int) {
	firstRock := 0

	for i := len(ch.tunnel) - 1; i >= 0; i-- {
		if strings.Contains(ch.tunnel[i], "#") || strings.Contains(ch.tunnel[i], "-") {
			firstRock = i
			break
		}
	}

	ch.tunnel = append(ch.tunnel[:firstRock+1], []string{".......", ".......", "......."}...)
	for i := len(ch.rocks[rockType]) - 1; i >= 0; i-- {
		ch.tunnel = append(ch.tunnel, ch.rocks[rockType][i])
	}

	for {
		// blow
		atLeftWall := false
		atRightWall := false
		b := ch.blows[ch.currentBlow%len(ch.blows)]
		ch.currentBlow++

		for i := len(ch.tunnel) - 1; i >= 0; i-- {
			if strings.Contains(ch.tunnel[i], "@") {
				if strings.HasPrefix(ch.tunnel[i], "@") || strings.Contains(ch.tunnel[i], "#@") {
					atLeftWall = true
				}

				if strings.HasSuffix(ch.tunnel[i], "@") || strings.Contains(ch.tunnel[i], "@#") {
					atRightWall = true
				}
			}
		}

		if b == '>' && !atRightWall {
			ch.move("right")
		} else if b == '<' && !atLeftWall {
			ch.move("left")
		}

		if ch.drop() {
			break
		}
	}
}

func newChamber(inpStr string) chamber {
	return chamber{
		tunnel: []string{"-------"},
		rocks:  createRockList(),
		blows:  []rune(inpStr),
	}
}

func Run(fName string) {
	fmt.Println("Starting Run (day 17)")

	inpBytes, _ := os.ReadFile(fName)
	inpStr := strings.TrimSpace(string(inpBytes))

	combiSeen := make(map[int]int)
	ch := newChamber(inpStr)
	heightPerBatch := 0
	batchSize := 0
	// height
	firstRepeater := 0
	for rockNo := 0; rockNo < 5000; rockNo++ {
		blowNo := ch.currentBlow % len(ch.blows)
		rockIndex := rockNo % len(ch.rocks)
		fmt.Printf("rockNo:%d (%d) and blowNo%d, (%d) \n", rockNo, rockIndex, ch.currentBlow, blowNo)
		prevRockNo, ok := combiSeen[blowNo]
		if ok && rockIndex == 0 {
			fmt.Printf("Repeating at rockNo:%d rockindex:%d and blowNo:%d\n", rockNo, rockIndex, blowNo)
			ch.print()
			batchSize = rockNo - prevRockNo
			firstRepeater = prevRockNo
			// heightPerBatch = len(ch.tunnel) - 4
			break
		}
		combiSeen[blowNo] = rockNo
		ch.dropNewRock(rockNo % len(ch.rocks))
	}

	fmt.Printf("firstRepeater @:%d, batchSize:%d\n", firstRepeater, batchSize)

	fmt.Printf("height:%d, roclls:%d \n", heightPerBatch, batchSize)
	height := heightPerBatch * (1000000000000 / batchSize)
	rest := 1000000000000 % batchSize

	fmt.Printf("h:%d, rest:%d\n", height, rest)
	freshChamber := newChamber(inpStr)
	freshChamber.currentBlow = ch.currentBlow
	for rockNo := 0; rockNo < rest; rockNo++ {
		freshChamber.dropNewRock(rockNo % len(freshChamber.rocks))
	}
	fmt.Printf("len:%d, answer:%d\n", len(freshChamber.tunnel), len(freshChamber.tunnel)-4+height)
}
