// Part 1

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
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
			m[currPoint] = string(c)

			if string(c) == PLAYER {
				start = currPoint
			}

			currPoint.x++
		}
		currPoint.x = 0
		currPoint.y++
	}

	return m
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	strings := make([]string, 0)

	for input.Scan() {
		line := input.Text()
		strings = append(strings, line)
	}

	m := parseMap(strings)
	fmt.Println(m)
	fmt.Println(start)
}
