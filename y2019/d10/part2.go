// Part 2

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"os/exec"
	"sort"
	"time"
)

var CENTER Point = Point{18, 20}

type Slope struct {
	rise, run float64
}

type Point struct {
	x, y int
}

type Points struct {
	points        []Point
	width, height int
}

func gcd(a, b float64) float64 {
	for b != 0 {
		t := b
		b = math.Mod(a, b)
		a = t
	}

	return math.Abs(a)
}

func getAsteroidPoints(scanner *bufio.Scanner) Points {
	points := make([]Point, 0)

	i := 0
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		for j, v := range line {
			if string(v) == "#" {
				points = append(
					points,
					Point{i, j},
				)
			}
			count++
		}
		i += 1
	}

	out := Points{points, count / i, i}
	return out
}

func getMostInView(points Points) map[Slope][]Point {
	slopes := make(map[Point]map[Slope][]Point, 0)

	for _, center := range points.points {
		for _, pt := range points.points {
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
					slopes[center][slope] = append(
						slopes[center][slope],
						pt,
					)
				} else {
					slopes[center] = make(map[Slope][]Point)
					slopes[center][slope] = []Point{pt}
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

func (pts Points) Contains(pt Point) bool {
	for _, p := range pts.points {
		if p == pt {
			return true
		}
	}

	return false
}

func (pts Points) Print() {
	output := ""
	for i := 0; i < pts.height; i++ {
		for j := 0; j < pts.width; j++ {
			pt := Point{i, j}
			if pts.Contains(pt) {
				if pt == CENTER {
					output += "X"
				} else {
					output += "#"
				}
			} else {
				output += "."
			}

		}
		output += "\n"
	}

	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	fmt.Println(output)
}

func dist(pt Point) float64 {
	a := math.Pow(float64(pt.x-CENTER.x), 2)
	b := math.Pow(float64(pt.y-CENTER.y), 2)

	return math.Sqrt(a + b)
}

func closestToCenter(pts []Point) Point {
	minDist := float64(math.MaxInt32)
	minIndex := 0

	for i, pt := range pts {
		d := dist(pt)
		if d < minDist {
			minDist = d
			minIndex = i
		}
	}

	return pts[minIndex]
}

func remove(pts []Point, removeMe Point) []Point {
	out := make([]Point, 0)

	for _, pt := range pts {
		if pt != removeMe {
			out = append(out, pt)
		}
	}

	return out
}

func remove200(points Points, slopes map[Slope][]Point) {
	keys := make([]Slope, 0)

	for k := range slopes {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		iRads := math.Atan2(keys[i].rise, keys[i].run)
		jRads := math.Atan2(keys[j].rise, keys[j].run)
		return iRads > jRads
	})

	i := 0
	count := 0
	for {
		points.Print()

		// Get the the points along the current angle
		p := keys[i]
		pts := slopes[p]

		// Get the closest point
		closest := closestToCenter(pts)

		// Remove the closest point
		slopes[p] = remove(pts, closest)
		points.points = remove(points.points, closest)

		fmt.Println(closest)
		time.Sleep(time.Second / 50.)
		i++
		count++

		if count == 200 {
			x := closest.y * 100
			y := closest.x
			fmt.Println(x + y)
			break
		}

		// Loop back around
		if i >= len(keys) {
			i = 0
		}
	}

}

func main() {
	input := bufio.NewScanner(os.Stdin)
	points := getAsteroidPoints(input)

	slopes := getMostInView(points)
	remove200(points, slopes)
}
