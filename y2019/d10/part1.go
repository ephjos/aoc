// Part 1

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X int // Distance from left
	Y int // Distance from top
}

func getAsteroidPoints(scanner *bufio.Scanner) []Point {
	points := make([]Point, 0)

	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, v := range line {
			if string(v) == "#" {
				points = append(points, Point{i, j})
			}
		}
		j += 1
	}

	return points
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	points := getAsteroidPoints(input)
	fmt.Println(points)
}
