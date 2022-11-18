// Part 1

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x, y int
}

type PointQueue struct {
	items []Point
}

func NewPointQueue() *PointQueue {
	return &PointQueue{[]Point{}}
}

func (pq *PointQueue) Enqueue(p *Point) {
	pq.items = append(pq.items, *p)
}

func (pq *PointQueue) Dequeue() *Point {
	item := pq.items[0]
	pq.items = pq.items[1:]
	return &item
}

func (pq *PointQueue) Front() Point {
	return pq.items[0]
}

func (pq *PointQueue) IsEmpty() bool {
	return len(pq.items) == 0
}

func (pq *PointQueue) Size() int {
	return len(pq.items)
}

type Warp struct {
	p1, p2 Point
}

type Map struct {
	view   map[Point]string
	player Point
	start  []Point
	end    []Point
}

const (
	OPEN   = "."
	WALL   = "#"
	EMPTY  = " "
	PLAYER = "@"

	U = 10
	D = 11
	L = 12
	R = 13
)

var DIRS = []int{U, D, L, R}

func IsCapitalLetter(s string) bool {
	runes := []rune(s)
	if len(runes) != 1 {
		return false
	}

	r := runes[0]

	return int(r) > 64 && int(r) < 91
}

func parseMap(lines []string) Map {
	point := Point{0, 0}
	start := make([]Point, 2)
	end := make([]Point, 2)
	player := Point{0, 0}
	viewMap := make(map[Point]string, 0)

	// Read view points in
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			char := string(line[i])

			switch char {
			case EMPTY:
				break
			case OPEN, WALL:
				fallthrough
			default:
				viewMap[point] = char
			}

			point.x++
		}
		point.x = 0
		point.y++
	}

	for p, s := range viewMap {
		if s == "A" {
			tp := Point{p.x, p.y}
			tp.y--

			v, ok := viewMap[tp]

			if ok && v == "A" {
				player = Point{p.x, p.y + 1}
				start[0] = tp
				start[1] = p
				continue
			}
			tp = Point{p.x, p.y}
			tp.x--

			v, ok = viewMap[tp]

			if ok && v == "A" {
				player = Point{p.x + 1, p.y}
				start[0] = tp
				start[1] = p
				continue
			}
		}

		if s == "Z" {
			tp := Point{p.x, p.y}
			tp.y--

			v, ok := viewMap[tp]

			if ok && v == "Z" {
				end[0] = tp
				end[1] = p
				continue
			}
			tp = Point{p.x, p.y}
			tp.x--

			v, ok = viewMap[tp]

			if ok && v == "Z" {
				end[0] = tp
				end[1] = p
				continue
			}
		}
	}

	m := Map{
		viewMap,
		player,
		start,
		end,
	}

	return m
}

func drawMap(m Map) {
	minX := math.MaxInt32
	minY := math.MaxInt32
	maxX := 0
	maxY := 0

	for p := range m.view {
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

	//size := 110

	out := "\x1b[2;0H"
	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			p := Point{j, i}
			ch := m.view[p]

			if ch == "" {
				ch = OPEN
			}

			if p == m.player {
				ch = PLAYER
			}

			clr := 40

			switch ch {
			case WALL:
				clr = 47
				ch = " "
			case OPEN:
				ch = " "
			case PLAYER:
				clr = 103
				ch = " "
			}

			for _, sp := range m.start {
				if p == sp {
					clr = 106
				}
			}

			for _, ep := range m.end {
				if p == ep {
					clr = 41
				}
			}

			ch += " "

			out += fmt.Sprintf("\x1b[%dm%s", clr, string(ch))
		}
		fmt.Println()
		out += fmt.Sprint("\x1b[0m\n")
	}

	fmt.Println(out)
}

func moveDir(point Point, dir int) Point {
	switch dir {
	case U:
		return Point{point.x, point.y - 1}
	case D:
		return Point{point.x, point.y + 1}
	case L:
		return Point{point.x - 1, point.y}
	case R:
		return Point{point.x + 1, point.y}
	}

	panic("Bad move")
}

func getPortal(m Map, point Point, dir int) string {
	switch dir {
	case U:
		p1 := Point{point.x, point.y - 1}
		s1 := m.view[p1]
		s2 := m.view[point]
		return s1 + s2
	case D:
		p1 := Point{point.x, point.y + 1}
		s1 := m.view[p1]
		s2 := m.view[point]
		return s2 + s1
	case L:
		p1 := Point{point.x - 1, point.y}
		s1 := m.view[p1]
		s2 := m.view[point]
		return s1 + s2
	case R:
		p1 := Point{point.x + 1, point.y}
		s1 := m.view[p1]
		s2 := m.view[point]
		return s2 + s1
	}

	panic("Couldn't get portal")
}

func jumpPortal(m Map, portal string, currPoint Point) Point {
	c1 := string(portal[0])
	c2 := string(portal[1])

	for p, v := range m.view {
		if p == currPoint {
			continue
		}

		if v == c1 {
			// Check below
			tp := Point{p.x, p.y}
			tp.y++

			if tp == currPoint {
				continue
			}

			if m.view[tp] == c2 {
				// Is open space to the top or bottom
				p1 := Point{tp.x, tp.y + 1}
				if m.view[p1] == OPEN {
					return p1
				}

				p2 := Point{p.x, p.y - 1}
				if m.view[p2] == OPEN {
					return p2
				}

			}

			// Check right
			tp = Point{p.x, p.y}
			tp.x++

			if tp == currPoint {
				continue
			}

			if m.view[tp] == c2 {
				// Is open space to the right or left
				p1 := Point{tp.x + 1, tp.y}
				if m.view[p1] == OPEN {
					return p1
				}

				p2 := Point{p.x - 1, p.y}
				if m.view[p2] == OPEN {
					return p2
				}

			}
		}
	}

	panic("Bad portal")
}

func BFS(m Map) int {

	// Initialize dist map, track distance to each
	// Point object
	dist := make(map[Point]int, 0)

	// Create initial start Point point
	start := m.player

	// Set start distance to 0
	dist[start] = 0

	// Create queue
	q := NewPointQueue()

	// Add start point to Queue
	q.Enqueue(&start)

	// While there are still items in the Queue
	for !q.IsEmpty() {
		// Get front of queue
		u := q.Dequeue()

		// Get known distance for current Point
		udist := dist[*u]

		/*
			m.player = *u
			drawMap(m)
			time.Sleep(time.Second / 500.)
		*/

		// For each neighbor (U,D,L,R)
		for _, dir := range DIRS {
			v := moveDir(*u, dir)
			// Get value at point in map
			char := m.view[v]

			// At start, no where to go
			if v == m.start[0] ||
				v == m.start[1] {
				continue
			}

			// If nb is a wall, do nothing
			if char == WALL ||
				char == EMPTY {
				continue
			}

			// At end, return
			if v == m.end[0] ||
				v == m.end[1] {
				return dist[*u]
			}

			if IsCapitalLetter(char) {
				portal := getPortal(m, v, dir)
				v = jumpPortal(m, portal, v)
			}

			// If we have a distance record already,
			// do nothing
			if dist[v] != 0 {
				continue
			}

			// Otherwise, update distance and add to queue
			dist[v] = udist + 1
			q.Enqueue(&v)
		}

	}

	return -1
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)

	for input.Scan() {
		line := input.Text()
		lines = append(lines, line)
	}

	m := parseMap(lines)
	dist := BFS(m)
	fmt.Println("Distance:", dist)
}
