// Part 1

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	EMPTY  = 0
	WALL   = 1
	BLOCK  = 2
	PADDLE = 3
	BALL   = 4
)

type Point struct {
	x, y int64
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	tokens := strings.Split(line, ",")

	ic := New(tokens)

	tileMap := make(map[Point]int64, 0)

	initialPos := Point{-1, -1}
	paddle := initialPos
	ball := initialPos

	score := int64(0)

	ic.Data[0] = 2

	inp := 0

	for ic.IsRunning {
		if paddle == initialPos || ball == initialPos {
			inp = 0
		} else {
			inp = int(math.Max(-1, math.Min(float64(ball.x-paddle.x), 1)))
		}

		x := ic.Run(inp)
		y := ic.Run(inp)
		z := ic.Run(inp)

		if x == -1 && y == 0 {
			score = z
		} else if z == PADDLE {
			paddle = Point{x, y}
		} else if z == BALL {
			ball = Point{x, y}
		} else {
			tileMap[Point{x, y}] = z
		}
	}

	fmt.Println(score)

}
