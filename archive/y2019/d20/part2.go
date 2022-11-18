// Part 1

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x, y, z int
}

func (p Point) Eq(o Point) bool {
	return (p.x == o.x) && (p.y == o.y)
}
func (p Point) Gt(o Point) bool {
	return (p.x > o.x) && (p.y > o.y)
}
func (p Point) Lt(o Point) bool {
	return (p.x < o.x) && (p.y < o.y)
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
	a, b             Point
	aInside, bInside bool
}

func (w *Warp) AddPoint(p, tl, br Point) {
	// Check if a is set
	if w.a.x == 0 {
		w.a = p
		w.aInside = p.Gt(tl) && p.Lt(br)
		return
	} else {
		// Set b
		w.b = p
		w.bInside = p.Gt(tl) && p.Lt(br)
		return
	}

	panic("Bad warp")
}

type Map struct {
	view   map[Point]string
	warps  map[string]Warp
	player Point
	start  Point
	end    Point
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
	point := Point{0, 0, 0}
	player := Point{0, 0, 0}
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

	minX := math.MaxInt32
	minY := math.MaxInt32
	maxX := 0
	maxY := 0

	for p, v := range viewMap {
		if v == WALL {
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
	}

	tl := Point{minX, minY, 0}
	br := Point{maxX, maxY, 0}

	warpMap := make(map[string]Warp, 0)

	for p, s := range viewMap {
		// We have a capital
		if IsCapitalLetter(s) {
			// Get the char's to the right and below
			n1 := Point{p.x, p.y + 1, 0} // Below
			c1 := viewMap[n1]
			n2 := Point{p.x + 1, p.y, 0} // Right
			c2 := viewMap[n2]

			// If below is also a letter
			if IsCapitalLetter(c1) {
				// Grab warp
				warpName := s + c1
				warp := warpMap[warpName]

				// Find where adjacent open point is (above or below)
				o1 := Point{p.x, p.y - 1, 0} // Above
				oc1 := viewMap[o1]

				// If above, set warp point to top letter
				if oc1 == OPEN {
					warp.AddPoint(p, tl, br)
				} else {
					// Else, set warp point to bottom letter
					warp.AddPoint(n1, tl, br)
				}

				warpMap[warpName] = warp
				continue
			}

			// Otherwise if right is also a letter
			if IsCapitalLetter(c2) {
				// Grab warp
				warpName := s + c2
				warp := warpMap[warpName]

				// Find where adjacent open point is (left or right)
				o1 := Point{p.x - 1, p.y, 0} // Left
				oc1 := viewMap[o1]

				// If above, set warp point to left letter
				if oc1 == OPEN {
					warp.AddPoint(p, tl, br)
				} else {
					// Else, set warp point to right letter
					warp.AddPoint(n2, tl, br)
				}

				warpMap[warpName] = warp
				continue
			}
		}
	}

	start := warpMap["AA"].a

	m := Map{
		viewMap,
		warpMap,
		player,
		start,
		warpMap["ZZ"].a,
	}

	player = NextOpenPoint(m, start)
	m.player = player

	return m
}

func NextOpenPoint(m Map, point Point) Point {
	for _, dir := range DIRS {
		p := moveDir(point, dir)

		if m.view[p] == OPEN {
			return p
		}
	}

	panic("No open neighbor")
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

	out := "\x1b[2;0H"
	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			p := Point{j, i, 0}
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
			/*
				for _, v := range m.warps {
					if p == v.a {
						if v.aInside {
							clr = 43
						} else {
							clr = 44
						}
					} else if p == v.b {
						if v.bInside {
							clr = 42
						} else {
							clr = 44
						}
					}
				}
			*/

			if p.Eq(m.start) {
				clr = 106
			}

			if p.Eq(m.end) {
				clr = 41
			}

			ch += " "

			out += fmt.Sprintf("\x1b[%dm%s", clr, string(ch))
		}
		out += fmt.Sprint("\x1b[0m\n")
	}

	fmt.Println(out)
}

func moveDir(point Point, dir int) Point {
	switch dir {
	case U:
		return Point{point.x, point.y - 1, 0}
	case D:
		return Point{point.x, point.y + 1, 0}
	case L:
		return Point{point.x - 1, point.y, 0}
	case R:
		return Point{point.x + 1, point.y, 0}
	}

	panic("Bad move")
}

func getPortal(m Map, point Point, dir int) string {
	switch dir {
	case U:
		p1 := Point{point.x, point.y - 1, 0}
		s1 := m.view[p1]
		s2 := m.view[point]
		return s1 + s2
	case D:
		p1 := Point{point.x, point.y + 1, 0}
		s1 := m.view[p1]
		s2 := m.view[point]
		return s2 + s1
	case L:
		p1 := Point{point.x - 1, point.y, 0}
		s1 := m.view[p1]
		s2 := m.view[point]
		return s1 + s2
	case R:
		p1 := Point{point.x + 1, point.y, 0}
		s1 := m.view[p1]
		s2 := m.view[point]
		return s2 + s1
	}

	panic("Couldn't get portal")
}

func jumpPortal(m Map, w Warp, currPoint Point) Point {
	z := currPoint.z

	if w.a.Eq(currPoint) {
		if w.aInside {
			z++
		} else {
			z--
		}

		p := NextOpenPoint(m, w.b)
		p.z = z
		return p
	} else if w.b.Eq(currPoint) {
		if w.bInside {
			z++
		} else {
			z--
		}

		p := NextOpenPoint(m, w.a)
		p.z = z
		return p
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
			time.Sleep(time.Second / 100.)
		*/

		// For each neighbor (U,D,L,R)
		for _, dir := range DIRS {
			v := moveDir(*u, dir)
			v.z = u.z

			// Get value at point in map
			t := v
			t.z = 0
			char := m.view[t]

			// At start, no where to go
			if v.Eq(m.start) {
				continue
			}

			// If nb is a wall, do nothing
			if char == WALL ||
				char == EMPTY {
				continue
			}

			// At end, return
			if v.Eq(m.end) {
				if v.z == 0 {
					return dist[*u]
				} else {
					continue
				}
			}

			contFlag := false
			// Check if on warp
			for _, w := range m.warps {
				if w.a.Eq(v) {
					if (v.z == 0 && w.aInside) ||
						(v.z != 0) {
						v = jumpPortal(m, w, v)
					} else {
						contFlag = true
					}
				} else if w.b.Eq(v) {
					if (v.z == 0 && w.bInside) ||
						(v.z != 0) {
						v = jumpPortal(m, w, v)
					} else {
						contFlag = true
					}
				}
			}

			if contFlag {
				continue
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
