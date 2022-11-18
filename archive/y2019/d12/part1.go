// Part 1

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x, y, z int
}

type Velocity = Position

type Moon struct {
	pos Position
	vel Velocity
}

func parseMoon(line string) Moon {
	line = strings.Trim(line, "<")
	line = strings.Trim(line, ">")
	tokens := strings.Split(line, ", ")

	xValue := strings.Split(tokens[0], "=")[1]
	yValue := strings.Split(tokens[1], "=")[1]
	zValue := strings.Split(tokens[2], "=")[1]

	x, _ := strconv.Atoi(xValue)
	y, _ := strconv.Atoi(yValue)
	z, _ := strconv.Atoi(zValue)

	p := Position{x, y, z}
	v := Velocity{0, 0, 0}
	m := Moon{p, v}

	return m
}

func getDiff(a, b int) int {
	if a == b {
		return 0
	} else if a < b {
		return 1
	} else if a > b {
		return -1
	}

	panic("Couldn't compare the values")
}

func updateVelocities(m1, m2 *Moon) {
	dx := getDiff(m1.pos.x, m2.pos.x)
	dy := getDiff(m1.pos.y, m2.pos.y)
	dz := getDiff(m1.pos.z, m2.pos.z)

	m1.vel.x += dx
	m2.vel.x += -dx

	m1.vel.y += dy
	m2.vel.y += -dy

	m1.vel.z += dz
	m2.vel.z += -dz
}

func updatePositions(moons *[]Moon) {
	for i := range *moons {
		moon := &(*moons)[i]
		moon.pos = Position{
			moon.pos.x + moon.vel.x,
			moon.pos.y + moon.vel.y,
			moon.pos.z + moon.vel.z,
		}
	}
}

func step(moons *[]Moon) {
	l := len(*moons)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			updateVelocities(&(*moons)[i], &(*moons)[j])
		}
	}

	updatePositions(moons)
}

func printMoon(m Moon) {
	fmt.Println(m.pos)
	fmt.Println(m.vel)
}

func printMoons(moons []Moon) {
	for _, moon := range moons {
		printMoon(moon)
	}
	fmt.Println()
}

func moonEnergy(moon Moon) int {
	xPE := math.Abs(float64(moon.pos.x))
	yPE := math.Abs(float64(moon.pos.y))
	zPE := math.Abs(float64(moon.pos.z))

	PE := xPE + yPE + zPE

	xK := math.Abs(float64(moon.vel.x))
	yK := math.Abs(float64(moon.vel.y))
	zK := math.Abs(float64(moon.vel.z))

	K := xK + yK + zK

	return int(PE * K)
}

func systemEnergy(moons []Moon) int {
	energy := 0

	for _, moon := range moons {
		energy += moonEnergy(moon)
	}

	return energy
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	moons := make([]Moon, 0)

	for input.Scan() {
		line := input.Text()
		moons = append(moons, parseMoon(line))
	}

	steps := 1000

	for i := 0; i < steps; i++ {
		step(&moons)
	}
	printMoons(moons)

	energy := systemEnergy(moons)
	fmt.Println("Total energy: ", energy)
}
