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
var memo = make(map[string]int, 0)

type Point struct {
	x, y int
}

type KeyInfo struct {
	point    Point
	name     string
	distance int
}

const (
	OPEN   = "."
	WALL   = "#"
	PLAYER = "@"
)

var start = Point{0, 0}

func parseMap(strings []string) map[Point]string {
	m := make(map[Point]string, 0)

	currPoint := Point{0, 0}

	for _, str := range strings {
		for _, c := range str {

			if string(c) == PLAYER {
				start = currPoint
			} else {
				m[currPoint] = string(c)
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
			ch = " " + ch + " "
			out += fmt.Sprintf("\x1b[%dm%s", clr, string(ch))

		}
		out += fmt.Sprint("\x1b[0m\n")
	}

	fmt.Println(out)
}

// def score(x, y, keys):
//   return min([score(key_x, key_y, new_key | keys) + steps_to(key_x, key_y)
//     for (key_x, key_y, new_key) in reachableCore(x, y, keys)])

// TODO: HASH PARAMS TO STRING
func hashScoreCall(point Point, keys []KeyInfo) string {
	return ""
}

// TODO: MEMOIZE
func score(point Point, keys []KeyInfo) int {
	scores := []int{}

	for _, k := range reachable(point, keys) {
		s := score(point, append(keys, k)) + k.distance
		scores = append(scores, s)
	}

	min := math.MaxInt32

	for _, s := range scores {
		if s < min {
			min = s
		}
	}

	return min
}

func reachable(point Point, keys []KeyInfo) []KeyInfo {
	return reachableCore(point, keys, make(map[Point]bool), 0)
}

func reachableCore(point Point, keys []KeyInfo, visited map[Point]bool, distance int) []KeyInfo {
	visited[point] = true

	neighbors := []Point{
		Point{point.x - 1, point.y},
		Point{point.x + 1, point.y},
		Point{point.x, point.y - 1},
		Point{point.x, point.y + 1},
	}

	out := make([]KeyInfo, 0)

	for _, neighbor := range neighbors {
		str := maze[neighbor]
		if visited[neighbor] != true {
			switch str {
			case WALL:
			case OPEN, "":
				// Recursive call
				sub := reachableCore(neighbor, keys, visited, distance+1)
				for _, v := range sub {
					out = append(out, v)
				}
			case strings.ToLower(str):
				// Add key
				out = append(out, KeyInfo{neighbor, str, distance})
				// Recursive call
				sub := reachableCore(neighbor, keys, visited, distance+1)
				for _, v := range sub {
					out = append(out, v)
				}
			case strings.ToUpper(str):
				// Handle door
				if Contains(keys, strings.ToLower(str)) {
					// We have the key
					// Recursive call
					sub := reachableCore(neighbor, keys, visited, distance+1)
					for _, v := range sub {
						out = append(out, v)
					}
				}
			}
		}
	}

	return out
}

func Contains(m []KeyInfo, str string) bool {
	for _, v := range m {
		if v.name == str {
			return true
		}
	}
	return false
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	strings := make([]string, 0)

	for input.Scan() {
		line := input.Text()
		strings = append(strings, line)
	}

	maze = parseMap(strings)

	var m []KeyInfo

	s := score(start, m)
	fmt.Println(s)
}
