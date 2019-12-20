// Part 1

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var maze = make(map[Point]string, 0)
var start = Point{0, 0}
var ALL_KEYS = make([]string, 0)

type Point struct {
	x, y int
}

type KeyInfo struct {
	point Point
	name  string
}

type KeySet struct {
	data []KeyInfo
}

func (k *KeySet) Contains(ki KeyInfo) bool {
	for _, v := range k.data {
		if v == ki {
			return true
		}
	}

	return false
}

func (k *KeySet) ContainsName(s string) bool {
	for _, v := range k.data {
		if v.name == s {
			return true
		}
	}

	return false
}

func (k *KeySet) Add(ki KeyInfo) {
	if !k.Contains(ki) {
		k.data = append(k.data, ki)
	}
}

func (k *KeySet) Combine(ki *KeySet) *KeySet {
	out := &KeySet{[]KeyInfo{}}

	for _, v := range k.data {
		out.Add(v)
	}
	for _, v := range ki.data {
		out.Add(v)
	}

	return out
}

func (k *KeySet) Copy() *KeySet {
	out := make([]KeyInfo, len(k.data))
	copy(out, k.data)

	return &KeySet{out}
}

const (
	OPEN   = "."
	WALL   = "#"
	PLAYER = "@"
)

func parseMap(strings []string) map[Point]string {
	m := make(map[Point]string, 0)

	currPoint := Point{0, 0}

	for _, str := range strings {
		for _, c := range str {

			if string(c) == PLAYER {
				start = currPoint
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

	return m
}

func print(m map[Point]string) {
	size := 81
	out := "\x1b[2;0H"

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			p := Point{j, i}

			clr := 40
			ch := m[p]

			if p == start {
				ch = PLAYER
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
					clr = 103
				}
			}
			ch = "" + ch + " "
			out += fmt.Sprintf("\x1b[%dm%s", clr, string(ch))

		}
		out += fmt.Sprint("\x1b[0m\n")
	}

	fmt.Println(out)
}

func score(point Point, ks *KeySet, visited map[Point]bool) int {
	if len(ALL_KEYS) == len(ks.data) {
		return 0
	}

	if visited[point] {
		return 0
	}

	visited[point] = true

	scores := make([]int, 0)

	for _, k := range reachable(point, ks, make(map[Point]bool)).data {
		nKs := ks.Copy()
		nKs.Add(k)
		s := score(k.point, nKs, visited) + 1 // + distance from point to k.point
		scores = append(scores, s)
	}

	min := math.MaxInt32

	for _, s := range scores {
		if s < min {
			min = s
		}
	}

	fmt.Println(point)
	fmt.Println(ks)
	fmt.Println(visited)
	fmt.Println(scores)
	fmt.Println()

	return min
}

func reachable(point Point, ks *KeySet, visited map[Point]bool) *KeySet {
	if visited[point] {
		return ks
	}

	visited[point] = true
	nbs := []Point{
		Point{point.x - 1, point.y},
		Point{point.x + 1, point.y},
		Point{point.x, point.y - 1},
		Point{point.x, point.y + 1},
	}

	out := ks.Copy()

	for _, nb := range nbs {
		v := maze[nb]

		switch v {
		case WALL:
		case OPEN:
			out = out.Combine(reachable(nb, ks, visited))
		case strings.ToLower(v):
			if v != "" && !ks.ContainsName(v) {
				ks.Add(KeyInfo{nb, v})
				out = out.Combine(reachable(nb, ks, visited))
			}
		case strings.ToUpper(v):
			if ks.ContainsName(strings.ToLower(v)) {
				out = out.Combine(reachable(nb, ks, visited))
			}
		}
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
	ks := &KeySet{[]KeyInfo{}}

	fmt.Println(score(start, ks, make(map[Point]bool)))
}
