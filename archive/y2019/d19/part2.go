// Part 1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	STATIONARY int = 0
	PULLED     int = 1
)

type Point struct {
	x, y int
}

func getClosestSquare(ic *IntCode, size int) int {

	x := 0
	y := 10
	for {
		fmt.Println(x, y)
		cp := ic.Copy()

		cp.AddInput(x)
		cp.AddInput(y)

		status := cp.Run()

		// Point is in beam
		if int(status) == PULLED {

			cp := ic.Copy()

			cp.AddInput(x + (size - 1))
			cp.AddInput(y - (size - 1))

			status := cp.Run()

			// Opposite point is in beam
			if int(status) == PULLED {
				return (x * 10000) + (y - (size - 1))
			}

			x = 0
			y++
			continue
		}
		x++
	}

	return 0
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	tokens := strings.Split(line, ",")

	ic := NewIntCode(tokens)

	fmt.Println(getClosestSquare(ic, 100))
}
