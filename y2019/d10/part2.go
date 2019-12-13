// Part 2

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Slope struct {
	rise, run float64
}

type Point struct {
	x, y int
}

func gcd(a, b float64) float64 {
	for b != 0 {
		t := b
		b = math.Mod(a, b)
		a = t
	}

	return math.Abs(a)
}

func getAsteroidPoints(scanner *bufio.Scanner) []Point {
	points := make([]Point, 0)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		for j, v := range line {
			if string(v) == "#" {
				points = append(
					points,
					Point{i, j},
				)
			}
		}
		i += 1
	}

	return points
}

func getMostInView(points []Point) map[Slope]int {
	slopes := make(map[Point]map[Slope]int, 0)

	for _, center := range points {
		for _, pt := range points {
			if pt != center {
				deltaY := float64(pt.y - center.y)
				deltaX := float64(pt.x - center.x)
				gcd := gcd(deltaY, deltaX)

				if gcd != 0 {
					deltaY /= gcd
					deltaX /= gcd
				}

				slope := Slope{deltaY, deltaX}

				if slopes[center] != nil {
					slopes[center][slope] += 1
				} else {
					slopes[center] = make(map[Slope]int, 0)
					slopes[center][slope] = 1
				}
			}
		}
	}

	max := 0
	maxIndex := Point{}
	for i, s := range slopes {
		count := len(s)
		if count > max {
			max = count
			maxIndex = i
		}
	}

	return slopes[maxIndex]
}

func printPoints(points []Point, slopes map[Slope]int) {
	for slope := range slopes {
		fmt.Println(slope)
		rads := math.Atan2(slope.rise, slope.run)
		degs := rads * (180. / math.Pi)

		fmt.Println(rads, degs)
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	points := getAsteroidPoints(input)

	slopes := getMostInView(points)
	printPoints(points, slopes)
}
