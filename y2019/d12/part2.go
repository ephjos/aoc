// Part 2

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

func gcd(a, b float64) float64 {
	for b != 0 {
		t := b
		b = math.Mod(a, b)
		a = t
	}

	return math.Abs(a)
}

func lcm(a, b int) int {
	x := float64(a)
	y := float64(b)
	return int(math.Floor((x * y) / gcd(x, y)))
}

func checkX(ms1, ms2 []Moon) bool {
	for i := range ms1 {
		m1 := ms1[i]
		m2 := ms2[i]

		if m1.pos.x != m2.pos.x || m1.vel.x != m2.vel.x {
			return false
		}
	}

	return true
}
func checkY(ms1, ms2 []Moon) bool {
	for i := range ms1 {
		m1 := ms1[i]
		m2 := ms2[i]

		if m1.pos.y != m2.pos.y || m1.vel.y != m2.vel.y {
			return false
		}
	}

	return true
}
func checkZ(ms1, ms2 []Moon) bool {
	for i := range ms1 {
		m1 := ms1[i]
		m2 := ms2[i]

		if m1.pos.z != m2.pos.z || m1.vel.z != m2.vel.z {
			return false
		}
	}

	return true
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	moons := make([]Moon, 0)

	for input.Scan() {
		line := input.Text()
		moons = append(moons, parseMoon(line))
	}

	orig := make([]Moon, len(moons))
	copy(orig, moons)

	count := 1

	xCount := 0
	yCount := 0
	zCount := 0

	xCheck := false
	yCheck := false
	zCheck := false

	for {
		step(&moons)

		if xCheck == yCheck && yCheck == zCheck && zCheck == true {
			fmt.Println(xCount, yCount, zCount)
			fmt.Println()
			break
		}

		if checkX(orig, moons) && !xCheck {
			xCheck = true
			xCount = count
		}
		if checkY(orig, moons) && !yCheck {
			yCheck = true
			yCount = count
		}
		if checkZ(orig, moons) && !zCheck {
			zCheck = true
			zCount = count
		}

		count++
	}

	fmt.Printf("Number: %d\n", lcm(xCount, lcm(yCount, zCount)))
}
