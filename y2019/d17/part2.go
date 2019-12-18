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

	COMMA = 44
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
	//out := ""
	for i := minY; i < maxY-1; i++ {
		for j := minX; j < maxX; j++ {
			p := Point{j, i}
			r := viewMap[p]

			clr := 40
			clr = clr
			ch := string(r) + ""
			ch = ch

			switch r {
			case SCAFFOLD:
				clr = 46
			case SPACE:
				clr = 40
			default:
				clr = 103
			}

			out += fmt.Sprintf("\x1b[%dm%s", clr, string(ch))
			//out += string(r)
		}
		out += fmt.Sprint("\x1b[0m\n")
		//out += "\n"
	}

	fmt.Println(out)
}

func traverseScaffold(ic *IntCode) {
	//o := 48
	main := "A,B,A,B,C,A,B,C,A,C"
	a := "R,6,L,10,R,8"
	b := "R,8,R,12,L,8,L,8"
	c := "L,10,R,6,R,6,L,8"
	stream := "y"
	inputs := []string{main, a, b, c, stream}

	ic.Data[0] = 2

	for _, input := range inputs {
		for _, char := range input {
			ic.AddInput(int(char))
		}
		ic.AddInput(NEWLINE)
	}

	last := -1
	var output int64
	answer := 0

	fmt.Print("\x1b[2;0H")
	for ic.IsRunning {
		output = ic.Run()
		if output > 300 {
			answer = output
		}

		clr := 40
		ch := string(output) + ""
		ch = "  "

		switch output {
		case SCAFFOLD:
			clr = 46
		case SPACE:
			clr = 40
		default:
			clr = 103
		}

		fmt.Print("\x1b[%dm%s", clr, string(ch))

		if output == NEWLINE && last == NEWLINE {
			fmt.Print("\x1b[0m\n")
			fmt.Print("\x1b[2;0H")
		}
		last = output
	}
	fmt.Print("\x1b[0m\n")

	fmt.Printf("Dust collected: %d\n", answer)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	tokens := strings.Split(line, ",")

	ic := NewIntCode(tokens)

	traverseScaffold(ic)

}
