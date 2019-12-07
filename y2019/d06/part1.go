// Part 1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	label    string
	depth    int
	children []*Node
}

func (n *Node) UpdateDepth(increment int) {
	n.depth += increment
	for _, child := range n.children {
		child.UpdateDepth(increment)
	}
	return
}

func (n *Node) AddChild(c *Node) {
	c.UpdateDepth(n.depth + 1)
	n.children = append(n.children, c)
	return
}

func (n *Node) PreOrder() {
	fmt.Println(n)
	for _, child := range n.children {
		child.PreOrder()
	}
	return
}

func SumTo(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func GetOrbitCount(n *Node) int {
	if n.children == nil {
		return n.depth
	} else {
		sum := n.depth
		for _, child := range n.children {
			sum += GetOrbitCount(child)
		}

		return sum
	}
}

func handleOrbits(orbits [][]string) {
	nodeMap := make(map[string]*Node)

	for _, orbitPair := range orbits {
		body := orbitPair[0]
		bodyNode := &Node{body, 0, nil}
		orbiter := orbitPair[1]
		orbiterNode := &Node{orbiter, 0, nil}

		if nodeMap[body] == nil {
			nodeMap[body] = bodyNode
		}

		if nodeMap[orbiter] == nil {
			nodeMap[orbiter] = orbiterNode
		}

		nodeMap[body].AddChild(nodeMap[orbiter])
	}

	COM := nodeMap["COM"]
	fmt.Println(GetOrbitCount(COM))
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
