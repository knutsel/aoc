package day07

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dominikbraun/graph"
)

type Bag struct {
	Name string
}

func Run(fName string) {
	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	nameHash := func(b Bag) string {
		return b.Name
	}
	g := graph.New(nameHash, graph.Directed())

	for _, l := range strings.Split(strings.TrimSpace(inpStr), "\n") {
		fields := strings.Fields(l)
		newNode := strings.Join(fields[0:2], "-")
		_ = g.AddVertex(Bag{Name: newNode})

		if fields[4] == "no" {
			continue
		}

		for i := 4; i < len(fields); i += 4 {
			w, err := strconv.Atoi(fields[i])
			if err != nil {
				panic(err)
			}

			containsName := strings.Join(fields[i+1:i+3], "-")
			_ = g.AddVertex(Bag{Name: containsName})
			_ = g.AddEdge(containsName, newNode, graph.EdgeWeight(w))
		}
	}

	colors := map[string]bool{} // not sure if it'll list doubles
	_ = graph.DFS(g, "shiny-gold", func(value string) bool {
		fmt.Println(value)
		colors[value] = true
		return false
	})

	fmt.Printf("P1: %d\n", len(colors)-1) // -1 for shiny-gold self
}
