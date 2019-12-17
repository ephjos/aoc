// Part 1

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
	for i := minY - 1; i <= maxY; i++ {
		for j := minX - 1; j <= maxX; j++ {
			p := Point{j, i}
			clr := 40
			ch := "  "

			r := viewMap[p]

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

func getAlignmentParameters(viewMap map[Point]rune) []int {
	output := make([]int, 0)

	for point, r := range viewMap {
		if r == SCAFFOLD {
			x := point.x
			y := point.y

			u := Point{x, y - 1}
			d := Point{x, y + 1}
			l := Point{x - 1, y}
			r := Point{x + 1, y}

			uv := viewMap[u]
			dv := viewMap[d]
			lv := viewMap[l]
			rv := viewMap[r]

			if uv == SCAFFOLD && uv == dv && dv == lv && lv == rv {
				output = append(output, x*y)
			}
		}
	}

	return output
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	tokens := strings.Split(line, ",")

	ic := NewIntCode(tokens)

	viewMap := parseView(ic)
	print(viewMap)
	aParams := getAlignmentParameters(viewMap)

	sum := 0
	for _, v := range aParams {
		sum += v
	}
	fmt.Println(sum)

}
