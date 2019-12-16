// Part 1

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

var visited = make(map[Point]bool, 0)
var flooded = make(map[Point]bool, 0)
var DIRECTIONS = []Direction{UP, DOWN, LEFT, RIGHT}
var START = Point{0, 0}
var DEST = Point{0, 0}

type Point struct {
	x, y int
}

type Tile = int
type Direction = int

const (
	WALL Tile = iota
	OPEN
	DESTINATION

	UP    Direction = 1
	DOWN  Direction = 2
	LEFT  Direction = 3
	RIGHT Direction = 4
)

type Node struct {
	parent *Node
}

func (n *Node) setParent(node *Node) {
	n.parent = node
}

func (p Point) AddDirection(dir Direction) Point {
	switch dir {
	case UP:
		return Point{p.x, p.y + 1}
	case DOWN:
		return Point{p.x, p.y - 1}
	case LEFT:
		return Point{p.x - 1, p.y}
	case RIGHT:
		return Point{p.x + 1, p.y}
	}

	return p
}

func printVisited() {
	minX := math.MaxInt32
	minY := math.MaxInt32

	maxX := math.MinInt32
	maxY := math.MinInt32

	for p := range visited {
		if p.x < minX {
			minX = p.x
		}

		if p.x > maxX {
			maxX = p.x
		}

		if p.y < minY {
			minY = p.y
		}

		if p.y > maxY {
			maxY = p.y
		}
	}

	out := ""
	for i := minY - 4; i <= maxY; i++ {
		for j := minX - 1; j <= maxX+5; j++ {
			p := Point{i, j}
			if visited[p] == true {
				if p == START {
					out += "S"
				} else if p == DEST {
					out += "D"
				} else {
					out += "x"
				}
			} else {
				out += " "
			}
		}
		out += "\n"
	}

	fmt.Println(out)
}

func printFlooded() {
	minX := math.MaxInt32
	minY := math.MaxInt32

	maxX := math.MinInt32
	maxY := math.MinInt32

	for p := range flooded {
		if p.x < minX {
			minX = p.x
		}

		if p.x > maxX {
			maxX = p.x
		}

		if p.y < minY {
			minY = p.y
		}

		if p.y > maxY {
			maxY = p.y
		}
	}

	out := ""
	for i := minY - 4; i <= maxY; i++ {
		for j := minX - 1; j <= maxX+5; j++ {
			p := Point{i, j}
			if flooded[p] == true {
				if p == START {
					out += "S "
				} else if p == DEST {
					out += "D "
				} else {
					out += "o "
				}
			} else {
				out += "  "
			}
		}
		out += "\n"
	}

	fmt.Println(out)
}

func shortestPathToControl(ic *IntCode, point Point, count int) {
	visited[point] = true

	for _, dir := range DIRECTIONS {
		cp := ic.Copy()
		out := cp.Run(dir)
		p := point.AddDirection(dir)

		switch int(out) {
		case WALL:
			break
		case OPEN:
			if visited[p] != true {
				shortestPathToControl(cp, p, count+1)
			}
			break
		case DESTINATION:
			DEST = point
			minutes := oxygenFlood(cp, p, 0)
			printFlooded()
			fmt.Printf("Minutes to flood: %d\n", minutes)
			return
		}
	}
}

func oxygenFlood(ic *IntCode, point Point, count int) int {
	flooded[point] = true
	time.Sleep(time.Second / 1000.)

	outputs := make([]int, 0)
	for _, dir := range DIRECTIONS {
		cp := ic.Copy()
		out := cp.Run(dir)
		p := point.AddDirection(dir)

		switch int(out) {
		case WALL:
			break
		case OPEN:
			if flooded[p] != true {
				outputs = append(outputs, oxygenFlood(cp, p, count+1))
			}
			break
		}

	}

	max := count

	for _, v := range outputs {
		if v > max {
			max = v
		}
	}

	return max
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	tokens := strings.Split(line, ",")

	ic := NewIntCode(tokens)

	shortestPathToControl(ic, START, 0)

}
