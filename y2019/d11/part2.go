// Part 1

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	BLACK int = 0
	WHITE int = 1
	LEFT  int = 0
	RIGHT int = 1

	U int = 0
	R int = 1
	D int = 2
	L int = 3
)

type Point struct {
	x, y int
}

func updateDirection(dir, turn int) int {
	if turn == LEFT {
		dir--
	} else if turn == RIGHT {
		dir++
	}

	if dir == 4 {
		dir = U
	}

	if dir == -1 {
		dir = L
	}

	return dir
}

func movePoint(point Point, dir int) Point {
	switch dir {
	case U:
		return Point{point.x, point.y + 1}
	case R:
		return Point{point.x + 1, point.y}
	case D:
		return Point{point.x, point.y - 1}
	case L:
		return Point{point.x - 1, point.y}
	}

	panic("Unknown direction, couldn't move point")
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	tokens := strings.Split(line, ",")

	ic := MakeIntCode(tokens)
	ch := make(chan int)
	go ic.Compute(ch)

	colorMap := make(map[Point]int, 0)

	// Initialize current point
	point := Point{0, 0}

	// Initialize direction
	dir := U

	// Initialize first color
	colorMap[point] = 1

	for {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		for i := 5; i > -15; i-- {
			for j := 0; j < 50; j++ {
				if colorMap[Point{j, i}] == 1 {
					fmt.Print("XX")
				} else {
					fmt.Print("  ")
				}
			}
			fmt.Println()
		}
		time.Sleep(50000000)

		ch <- colorMap[point]
		paintColor := <-ch
		turnDirection := <-ch

		colorMap[point] = paintColor
		dir = updateDirection(dir, turnDirection)

		point = movePoint(point, dir)

	}

}
