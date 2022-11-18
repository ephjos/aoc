// Part 1

package main

import (
	"bufio"
	"fmt"
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
	x, y int
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	tokens := strings.Split(line, ",")

	ic := MakeIntCode(tokens)
	ch := make(chan int)
	go ic.Compute(ch)

	tileMap := make(map[Point]int, 0)

	for x := range ch {
		y := <-ch
		z := <-ch

		tileMap[Point{x, y}] = z
	}

	count := 0
	for _, id := range tileMap {
		if id == BLOCK {
			count++
		}
	}

	fmt.Println(count)

}
