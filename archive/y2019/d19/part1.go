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

func countBeamLocations(ic *IntCode, size int) int {
	points := make([]Point, 0)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			points = append(points, Point{j, i})
		}
	}

	count := 0
	viewMap := make(map[Point]string, 0)

	for _, point := range points {
		cp := ic.Copy()

		cp.AddInput(point.x)
		cp.AddInput(point.y)

		status := cp.Run()
		if int(status) == PULLED {
			count++
			viewMap[point] = "#"
		} else {
			viewMap[point] = "."
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(viewMap[Point{j, i}])
		}
		fmt.Println()
	}

	return count
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	tokens := strings.Split(line, ",")

	ic := NewIntCode(tokens)

	fmt.Println(countBeamLocations(ic, 50))
}
