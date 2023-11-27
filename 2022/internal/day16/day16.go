package day16

import (
	"fmt"
	"os"
	"strings"
)

type vertex struct {
	// name        string
	connections []*edge
	flowRate    int
	status      bool // true = open
	discovered  bool
}

type edge struct {
	toName string
	weight int
}

type graph map[string]*vertex

func (g graph) AddVertex(name string, flowRate int) {
	if _, ok := g[name]; !ok {
		g[name] = &vertex{flowRate: flowRate, connections: make([]*edge, 0), status: false}
	}
}

func (v *vertex) AddEdge(to string, weight int) {
	v.connections = append(v.connections, &edge{to, weight})
}

// type graph struct {
// 	nodes map[string][]edge
// }

// func newGraph() *graph {edges    []*edge
// 	return &graph{nodes: make(map[string][]edge)}
// }

// func (g *graph) addEdge(origin, destiny string, weight int) {
// 	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
// 	g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
// }

// func (g *graph) getEdges(node string) []edge {
// 	return g.nodes[node]
// }

func (g graph) print() {
	for n, v := range g {
		for _, c := range v.connections {
			fmt.Printf("v: %s status:%v, flowRate:%d, connected to %s w weight:%d\n", n, v.status, v.flowRate, c.toName, c.weight)
		}
	}
}

// func enqueue(queue []*vertex, element *vertex) []*vertex {
// 	queue = append(queue, element) // Simply append to enqueue.
// 	fmt.Println("Enqueued:", element)
// 	return queue
// }

// func dequeue(queue []*vertex) (*vertex, []*vertex) {
// 	element := queue[0] // The first element is the one to be dequeued.
// 	fmt.Println("Dequeued:", element)
// 	return element, queue[1:] // Slice off the element once it is dequeued.
// }

func bfsDistance(g graph, root string) map[string]int {
	dist := make(map[string]int)
	q := make([]string, 0)
	q = append(q, root)
	dist[root] = 0

	for {
		if len(q) == 0 {
			break
		}

		vName := q[0]
		q = q[1:]
		g[vName].discovered = true

		for _, adj := range g[vName].connections {
			if !g[adj.toName].discovered {
				dist[adj.toName] = dist[vName] + 1
				q = append(q, adj.toName)
			}
		}
	}

	return dist
}

func goDeeper(g graph, v string, timeLeft int, opened map[string]bool) int {
	if timeLeft <= 0 {
		return 0
	}
	timeLeft--
	// distances := bfsDistance(g, v)

	myFlow := 0
	if _, ok := opened[v]; !ok {
		myFlow = g[v].flowRate * timeLeft
		opened[v] = true
	}

	addedFlow := 0
	for _, e := range g[v].connections {
		addedFlow = goDeeper(g, e.toName, timeLeft, opened)
	}

	return myFlow + addedFlow
}

func dfs(g graph, v string, timeLeft int) {
	fmt.Printf("Discovered %s\n", v)

	g[v].discovered = true
	for _, e := range g[v].connections {
		if !g[e.toName].discovered {
			dfs(g, e.toName, timeLeft)
		}
	}
}

func bruteForce(g graph, v string, path string, flowByPath map[string]int, timeLeft int) {
	fmt.Printf("**Enter with %s and %d\n", path, timeLeft)
	if timeLeft <= 0 {
		fmt.Printf(">>> %+v timeLeft:%d\n", flowByPath, timeLeft)
		return
	}
	timeLeft--

	myFlow := 0
	if !strings.Contains(path, v) {
		myFlow = g[v].flowRate * timeLeft
		g[v].discovered = true
	}

	flowByPath[path] += myFlow

	for _, e := range g[v].connections {
		if !g[e.toName].discovered {
			bruteForce(g, e.toName, path+"."+e.toName, flowByPath, timeLeft)
		}
	}

	fmt.Printf("<<< %+v timeLeft:%d\n", flowByPath, timeLeft)
}

func Run(fName string) {
	fmt.Println("Starting Run (day 16)")

	inpBytes, _ := os.ReadFile(fName)
	inpStr := string(inpBytes)
	inputLines := strings.Split(strings.TrimSpace(inpStr), "\n")
	g := graph{}

	for _, l := range inputLines {
		var flowRate int64

		words := strings.Fields(l)
		originVertex := words[1]

		_, err := fmt.Sscanf(words[4], "rate=%d;", &flowRate)
		if err != nil {
			panic(err)
		}

		g.AddVertex(originVertex, int(flowRate))

		for _, t := range words[9:] {
			toNode := strings.Trim(t, ",")
			g[originVertex].AddEdge(toNode, 1) // weight is 1... Need to think
		}
	}

	// g.orderEdgesByFlowRate()
	// g.print()
	// for name, _ := range g {
	// 	sort.Slice(g[name].connections, func(i, j int) bool {
	// 		return g[g[name].connections[i].toName].flowRate > g[g[name].connections[j].toName].flowRate
	// 	})
	// }
	fmt.Println()
	g.print()

	// node := "AA"
	// minute := 30

	// distances := bfsDistance(g, node)
	// fmt.Printf("%+v\n", distances)
	// fmt.Printf("Minute %d node %s flow rate at %d\n", minute, node, g[node].flowRate)
	// maxRelease := 0
	// // opened := make(map[string]bool)
	// if g[node].flowRate == 0 {
	// 	for k, v := range distances {
	// 		wouldRelease := (minute - distances[k]) * g[k].flowRate
	// 		// maxRelease = max(wouldRelease, maxRelease)
	// 		fmt.Printf("start %s choice %s at  dist:%d, flow:%d redux -> %d max: %d\n", node, k, v, g[k].flowRate, wouldRelease, maxRelease)
	// 		// maxRelease = max(maxRelease, goDeeper(g,))
	// 	}
	// }
	flowByPath := make(map[string]int)
	bruteForce(g, "AA", "AA", flowByPath, 30)
	// fmt.Printf(">>> %d\n", m)
	// dfs(g, "AA")
}
