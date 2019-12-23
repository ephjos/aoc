// Part 2

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var maze = make(map[Point]string, 0)
var startPoints = []Point{
	Point{39, 39},
	Point{39, 41},
	Point{41, 41},
	Point{41, 39},
}
var ALL_KEYS = make([]string, 0)

type Point struct {
	x, y int
}

const (
	OPEN   = "."
	WALL   = "#"
	PLAYER = "@"
)

type XYKeys struct {
	point Point
	keys  []string
}

func NewXYKeys(point Point) *XYKeys {
	return &XYKeys{point, []string{}}
}

func (x *XYKeys) ContainsKey(k string) bool {
	for _, v := range x.keys {
		if v == k {
			return true
		}
	}

	return false
}

func (x *XYKeys) AddKey(k string) {
	if !x.ContainsKey(k) {
		x.keys = append(x.keys, k)
	}

	sort.Strings(x.keys)

}

func (x *XYKeys) Copy() *XYKeys {
	out := NewXYKeys(x.point)
	out.keys = make([]string, len(x.keys))
	copy(out.keys, x.keys)
	return out
}

func (x *XYKeys) Hash() string {
	out := ""
	out += strconv.Itoa(x.point.x) + ","
	out += strconv.Itoa(x.point.y) + ","

	for _, v := range x.keys {
		out += v + ","
	}

	return out
}

type XYKeysQueue struct {
	items []XYKeys
}

func NewXYKeysQueue() *XYKeysQueue {
	return &XYKeysQueue{[]XYKeys{}}
}

func (xq *XYKeysQueue) Enqueue(x *XYKeys) {
	xq.items = append(xq.items, *x)
}

func (xq *XYKeysQueue) EnqueueS(x *XYKeys, dist map[string]int) {
	xq.items = append(xq.items, *x)

	sort.Slice(xq.items, func(i, j int) bool {
		d1 := dist[xq.items[i].Hash()]
		d2 := dist[xq.items[j].Hash()]
		return d1 < d2
	})
}

func (xq *XYKeysQueue) Dequeue() *XYKeys {
	item := xq.items[0]
	xq.items = xq.items[1:]
	return &item
}

func (xq *XYKeysQueue) Front() XYKeys {
	return xq.items[0]
}

func (xq *XYKeysQueue) IsEmpty() bool {
	return len(xq.items) == 0
}

func (xq *XYKeysQueue) Size() int {
	return len(xq.items)
}

func StringContains(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}

	return false
}

func parseMap(strings []string) map[Point]string {
	m := make(map[Point]string, 0)

	currPoint := Point{0, 0}

	for _, str := range strings {
		for _, c := range str {

			if string(c) == PLAYER {
				m[currPoint] = OPEN
			} else {
				m[currPoint] = string(c)
				if string(c) != WALL && string(c) != OPEN &&
					(int(c) > 96 && int(c) < 123) {
					ALL_KEYS = append(ALL_KEYS, string(c))
				}
			}

			currPoint.x++
		}
		currPoint.x = 0
		currPoint.y++
	}

	sort.Strings(ALL_KEYS)

	return m
}

func print(m map[Point]string, current *XYKeys) {
	size := 81
	//size := 20
	out := "\x1b[2;0H"

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			p := Point{j, i}

			clr := 40
			ch := m[p]

			if p == current.point {
				ch = PLAYER
			}

			if current.ContainsKey(ch) {
				ch = OPEN
			}

			switch ch {
			case WALL:
				clr = 47
				ch = " "
			case PLAYER:
				clr = 106
			case OPEN:
				ch = " "
			default:
				if ch != "" {
					clr = 104
				}
			}
			ch = "" + ch + " "
			out += fmt.Sprintf("\x1b[%dm%s", clr, string(ch))

		}
		out += fmt.Sprint("\x1b[0m\n")
	}

	fmt.Println(out)
}

func GetNeighbors(point Point) []Point {
	return []Point{
		Point{point.x - 1, point.y},
		Point{point.x + 1, point.y},
		Point{point.x, point.y - 1},
		Point{point.x, point.y + 1},
	}
}

func BFS(startPoint Point, allKeys []XYKeys) (Point, int) {

	// Initialize dist map, track distance to each
	// XYKeys object
	dist := make(map[string]int, 0)

	// Create initial start XYKeys point
	start := NewXYKeys(startPoint)

	// Set start distance to 0
	dist[start.Hash()] = 0

	// Create queue
	q := NewXYKeysQueue()

	// Add start point to Queue
	q.Enqueue(start)

	// While there are still items in the Queue
	for !q.IsEmpty() {
		// Get front of queue
		u := q.Dequeue()

		// Get the current keys
		ukeys := u.keys

		// If we have all of the keys,
		// return the current point and distance to it
		if len(ukeys) == len(allKeys) {
			return u.point, dist[u.Hash()]
		}

		// Get known distance for current XYKeys
		udist := dist[u.Hash()]

		//print(maze, u)
		//fmt.Println("\x1b[2;0H")
		/*
			fmt.Println("All keys:", ALL_KEYS)
			fmt.Println("Current point + keys:", u)
			fmt.Println("Current distance:", udist)
		*/
		//fmt.Println("\x1b[0m\n")
		//time.Sleep(time.Second / 100.)

		// For each neighbor (U,D,L,R)
		for _, nb := range GetNeighbors(u.point) {
			// Get value at point in map
			char := maze[nb]

			// Create XYKeys at nb point
			v := u.Copy()
			v.point = nb

			// If nb is a wall, do nothing
			if char == WALL {
				continue
			}

			// If nb is on a lowercase letter (key),
			// try to add key to v
			if char == strings.ToLower(char) &&
				char != OPEN {
				v.AddKey(char)
			}

			// If nb is on an uppercase letter (door)
			if char == strings.ToUpper(char) &&
				char != OPEN {
				// If we don't have the key,
				// do nothing
				if !v.ContainsKey(strings.ToLower(char)) {
					continue
				}
			}

			// If we have a distance record already,
			// do nothing
			if dist[v.Hash()] != 0 {
				continue
			}

			// Otherwise, update distance and add to queue
			dist[v.Hash()] = udist + 1
			//q.EnqueueS(v, dist)
			q.Enqueue(v)
		}

	}

	return startPoint, -1
}

func reachable(point Point, visited map[Point]bool) []XYKeys {
	keys := make([]XYKeys, 0)
	if visited[point] == true {
		return keys
	}

	visited[point] = true

	for _, nb := range GetNeighbors(point) {
		char := maze[nb]

		if char == WALL {
			continue
		}

		if char != OPEN {
			if char == strings.ToUpper(char) {
				maze[nb] = OPEN
				keys = append(keys, reachable(nb, visited)...)
			}

			if char == strings.ToLower(char) {
				xyk := NewXYKeys(nb)
				xyk.AddKey(char)
				keys = append(keys, *xyk)
				keys = append(keys, reachable(nb, visited)...)
			}
		} else {
			keys = append(keys, reachable(nb, visited)...)
		}
	}

	out := make([]XYKeys, 0)

outerLoop:
	for _, k1 := range keys {
		for _, k2 := range out {
			if k1.point == k2.point {
				continue outerLoop
			}
		}

		out = append(out, k1)
	}

	return out
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	strings := make([]string, 0)

	for input.Scan() {
		line := input.Text()
		strings = append(strings, line)
	}

	maze = parseMap(strings)

	fmt.Println()

	keys := make([][]XYKeys, 4)
	for i, point := range startPoints {
		out := reachable(point, make(map[Point]bool))
		keys[i] = out
	}

	total := 0
	for i, point := range startPoints {
		_, dist := BFS(point, keys[i])
		total += dist
	}

	fmt.Println(total)

}
