// Part 2

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

const (
	SCAFFOLD = 35
	SPACE    = 46
	NEWLINE  = 10
)

func print(viewMap map[Point]rune) {
	//time.Sleep(time.Second / 150.)
	minX := math.MaxInt32
	minY := math.MaxInt32

	maxX := math.MinInt32
	maxY := math.MinInt32

	for p := range viewMap {
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
	for i := minY; i < maxY-1; i++ {
		for j := minX; j < maxX; j++ {
			p := Point{j, i}
			r := viewMap[p]

			clr := 40
			ch := "  "

			switch r {
			case SCAFFOLD:
				clr = 46
			case SPACE:
				clr = 40
			default:
				clr = 103
			}

			out += fmt.Sprintf("\x1b[%dm%s", clr, string(ch))
		}
		out += fmt.Sprint("\x1b[0m\n")
	}

	fmt.Println(out)
}

func parseView(ic *IntCode) map[Point]rune {
	var viewMap = make(map[Point]rune, 0)

	var point = Point{0, 0}

	for ic.IsRunning {
		out := ic.Run(0)

		viewMap[point] = rune(out)

		if out == NEWLINE {
			point.x = 0
			point.y++
		} else {
			point.x++
		}
	}

	return viewMap
}

func traverseScaffold(viewMap map[Point]rune) {
	o := 48
	main := []rune{'A', 'B', 'A', 'B', 'C', 'A', 'B', 'C', 'A', 'C'}
	a := []rune{'R', rune(6 + o), 'L', rune(10 + o), 'R', rune(8 + o)}
	b := []rune{'R', rune(8 + o), 'R', rune(12 + o), 'L', rune(8 + o), 'L', rune(8 + o)}
	c := []rune{'L', rune(10 + o), 'R', rune(6 + o), 'R', rune(6 + o), 'L', rune(8 + o)}

	fmt.Println(main)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	tokens := strings.Split(line, ",")

	ic := NewIntCode(tokens)

	viewMap := parseView(ic)
	print(viewMap)
	traverseScaffold(viewMap)

}
