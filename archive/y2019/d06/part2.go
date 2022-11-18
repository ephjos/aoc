// Part 2
// Instead of treating the relation as a tree,
// treat it as a graph and do Dijkstra's

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Node struct {
	Label       string
	Connections []*Node
}

func makeEdge(a, b *Node) {
	a.Connections = append(a.Connections, b)
	b.Connections = append(b.Connections, a)
}

type Vertex struct {
	Dist int
	Prev *Vertex
}

func GetMin(Q map[string]*Vertex) string {
	min := math.MaxInt32
	minLabel := ""

	for label, v := range Q {
		if v.Dist < min {
			min = v.Dist
			minLabel = label
		}
	}

	return minLabel
}

func Dijkstra(graph map[string]*Node, source, target string) int {
	Q := make(map[string]*Vertex, 0)

	for label := range graph {
		v := Vertex{math.MaxInt32, nil}
		Q[label] = &v
	}

	Q[source].Dist = 0

	for len(Q) > 0 {
		u := GetMin(Q)

		if u == target {
			return Q[u].Dist
		}

		neighbors := graph[u].Connections

		for _, neighbor := range neighbors {
			v := neighbor.Label
			if Q[v] != nil { // Still in Q
				alt := Q[u].Dist + 1 // All connection 0 weighted
				if alt < Q[v].Dist {
					Q[v].Dist = alt
					Q[v].Prev = Q[u]
				}
			}
		}

		delete(Q, u)
	}

	panic(fmt.Sprintf("Couldn't find %s\n", target))
}

func handleOrbits(orbits [][]string) {
	nodes := make(map[string]*Node)

	for _, orbitPair := range orbits {
		body := orbitPair[0]
		bodyNode := &Node{body, nil}
		orbiter := orbitPair[1]
		orbiterNode := &Node{orbiter, nil}

		if nodes[body] == nil {
			nodes[body] = bodyNode
		}

		if nodes[orbiter] == nil {
			nodes[orbiter] = orbiterNode
		}

		makeEdge(nodes[body], nodes[orbiter])

	}

	fmt.Println("Running Dijkstra ...")
	dist := Dijkstra(nodes, "SAN", "YOU")

	fmt.Println("Shortest path:", dist-2)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	var orbits [][]string

	for input.Scan() {
		line := input.Text()
		tokens := strings.Split(line, ")")
		orbits = append(orbits, tokens)
	}

	handleOrbits(orbits)
}
