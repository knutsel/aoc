package day19

import (
	"fmt"
	"os"
	"strings"
)

type cost struct {
	ore  int
	clay int
	obs  int
}

type blueprint struct {
	id   int
	cost map[string]cost
	// maxOreNeeded  int
	// maxClayNeeded int
	// maxObsNeeded  int
}

type State struct {
	// decisonTree map[string]string
	timeLeft int
	// cracked  int
	minerals map[string]int
	robots   map[string]int
}

func (s State) deepCopy() State {
	newState := State{
		timeLeft: s.timeLeft,
		minerals: map[string]int{},
		robots:   map[string]int{},
	}
	for k, v := range s.minerals {
		newState.minerals[k] = v
	}
	for k, v := range s.robots {
		newState.robots[k] = v
	}
	return newState
}

func (bp blueprint) maxGeodesCracked(inState State, path string, pathResult map[string]int) int {
	// fmt.Printf("Enter: %s\n", path)
	s := inState.deepCopy()
	if s.timeLeft <= 0 {
		fmt.Printf("timeLeft exit: %s result: %d\n", path, s.minerals["geo"])
		pathResult[path] = s.minerals["geo"]
		// return s.minerals["geo"]
		return s.minerals["geo"]
	}

	// if val, ok := pathResult[path]; ok {
	// 	return val
	// }

	// do the work
	fmt.Printf("state:%+v\n", s)
	for robo, num := range s.robots {
		for i := 0; i < num; i++ {
			// fmt.Printf("mining %s\n", robo)
			s.minerals[robo]++
		}
	}

	// curMax := s.minerals["geo"]
	pathResult[path] = s.minerals["geo"]

	choice := ""
	if bp.cost["geo"].ore <= s.minerals["ore"] && bp.cost["geo"].obs <= s.minerals["obs"] {
		s.robots["geo"]++
		s.minerals["ore"] -= bp.cost["geo"].ore
		s.minerals["obs"] -= bp.cost["geo"].obs
		choice = "geo"
	}

	if bp.cost["obs"].ore <= s.minerals["ore"] && bp.cost["obs"].clay <= s.minerals["clay"] && s.minerals["obs"] < bp.cost["geo"].obs {
		s.robots["obs"]++
		s.minerals["ore"] -= bp.cost["obs"].ore
		s.minerals["clay"] -= bp.cost["obs"].clay
		choice = "obs"
	}

	if bp.cost["clay"].ore <= s.minerals["ore"] && s.minerals["clay"] < bp.cost["obs"].clay {
		s.robots["clay"]++
		s.minerals["ore"] -= bp.cost["clay"].ore
		choice = "clay"
	}

	if bp.cost["ore"].ore <= s.minerals["ore"] {
		s.robots["ore"]++
		s.minerals["ore"] -= bp.cost["ore"].ore
		choice = "ore"
	}

	// fmt.Printf("bought %q\n", choice)
	s.timeLeft--
	return bp.maxGeodesCracked(s, path+"."+choice, pathResult)
	// curMax = max(result, curMax)

	// return curMax
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := strings.TrimSpace(string(inpBytes))

	blueprints := make([]blueprint, 0)

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		bp := blueprint{}
		// nolint: lll
		var a, b, c, d, e, f int
		_, err := fmt.Sscanf(l, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&bp.id, &a, &b, &c, &d, &e, &f)
		if err != nil {
			panic(err)
		}

		bp.cost = make(map[string]cost)
		bp.cost["ore"] = cost{ore: a}
		bp.cost["clay"] = cost{ore: b}
		bp.cost["obs"] = cost{ore: c, clay: d}
		bp.cost["geo"] = cost{ore: e, obs: f}

		// for _, v := range bp.cost {
		// 	bp.maxOreNeeded = max(bp.maxOreNeeded, v.ore)
		// 	bp.maxClayNeeded = max(bp.maxClayNeeded, v.clay)
		// 	bp.maxObsNeeded = max(bp.maxObsNeeded, v.obs)
		// }

		blueprints = append(blueprints, bp)
	}

	for i, bp := range blueprints {

		fmt.Printf("bp: %+v\n", bp)
		s := State{
			// decisonTree: map[string]string{},
			timeLeft: 24,
			// cracked:  0,
			minerals: map[string]int{"ore": 0, "clay": 0, "obs": 0, "geo": 0},
			robots:   map[string]int{"ore": 1, "clay": 0, "obs": 0, "geo": 0},
		}
		pathRes := make(map[string]int)
		gMax := bp.maxGeodesCracked(s, "", pathRes)

		fmt.Printf("%d pr:%+v -> ,max:%d\n", i, s, gMax)
		// break // TODO remove
	}
	// fmt.Printf("bp")
	//fmt.Printf("file:%s part1: %d part2:%d \n", fName, totalExposed, totalExposedExclPockets)
}
