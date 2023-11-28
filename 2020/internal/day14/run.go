package day14

import (
	"fmt"
	"os"
	"strings"
)

type mask struct {
	orMask  int64
	andMask int64
}

type computer struct {
	currentMask   mask
	memory        map[int64]int64
	currentMaskP2 string
	memoryP2      map[int64]int64
}

func (c *computer) setMask(l string) {
	maskStr := strings.TrimSpace(strings.Split(l, "=")[1])
	orString := strings.ReplaceAll(maskStr, "X", "0")
	andStr := strings.ReplaceAll(maskStr, "X", "1")
	c.currentMaskP2 = maskStr

	_, err := fmt.Sscanf(orString, "%b", &c.currentMask.orMask)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Sscanf(andStr, "%b", &c.currentMask.andMask)
	if err != nil {
		panic(err)
	}
}

func (c *computer) setMem(l string) {
	location := int64(0)
	value := int64(0)

	_, err := fmt.Sscanf(l, "mem[%d] = %d", &location, &value)
	if err != nil {
		panic(err)
	}

	newVal := value | c.currentMask.orMask
	newVal &= c.currentMask.andMask

	c.memory[location] = newVal

	// part2
	xLoc := []int{}
	addresses := []int64{}
	powerOfTwo := 0
	baseAddress := location

	for i := len(c.currentMaskP2) - 1; i >= 0; i-- {
		if c.currentMaskP2[i] == '1' {
			baseAddress |= 1 << powerOfTwo
		} else if c.currentMaskP2[i] == 'X' {
			xLoc = append(xLoc, powerOfTwo)
		}
		powerOfTwo++
	}

	addresses = append(addresses, baseAddress)

	for _, xIndex := range xLoc {
		for _, val := range addresses {
			with1 := val | 1<<xIndex
			with0 := with1 ^ 1<<xIndex
			addresses = append(addresses, with1, with0)
		}
	}

	for _, a := range addresses {
		c.memoryP2[a] = value
	}
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)

	c := computer{
		currentMask: mask{},
		memory:      map[int64]int64{},
		memoryP2:    map[int64]int64{},
	}

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		if strings.HasPrefix(l, "mask") {
			c.setMask(l)
		} else {
			c.setMem(l)
		}
	}

	sum := int64(0)
	sumP2 := int64(0)

	for _, v := range c.memory {
		sum += v
	}

	for _, v := range c.memoryP2 {
		sumP2 += v
	}

	fmt.Printf("P1:%d P2:%d\n", sum, sumP2)
}
