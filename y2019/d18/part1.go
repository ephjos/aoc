// Part 1

package main

import (
	"bufio"
	"fmt"
	"os"
)

var maze = make(map[Point]string, 0)
var memo = make(map[string]int, 0)
var start = Point{0, 0}

type Point struct {
	x, y int
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

func main() {
	input := bufio.NewScanner(os.Stdin)
	strings := make([]string, 0)

	for input.Scan() {
		line := input.Text()
		strings = append(strings, line)
	}

	maze = parseMap(strings)
	print(maze)
}
