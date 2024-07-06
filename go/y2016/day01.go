package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type day01 struct{}

func (_d day01) a(input string) string {
	type Point struct {
		x int
		y int
	}

	const (
		NORTH = iota
		EAST
		SOUTH
		WEST
	)

	curr := Point{0, 0}
	facing := NORTH

	for _, instr := range strings.Split(strings.Trim(input, " \n"), ", ") {
		dir := instr[0]
		amount, err := strconv.Atoi(instr[1:])
		if err != nil {
			panic(err)
		}

		switch dir {
		case 'L':
			switch facing {
			case NORTH:
				facing = WEST
			case WEST:
				facing = SOUTH
			case SOUTH:
				facing = EAST
			case EAST:
				facing = NORTH
			}
		case 'R':
			switch facing {
			case NORTH:
				facing = EAST
			case WEST:
				facing = NORTH
			case SOUTH:
				facing = WEST
			case EAST:
				facing = SOUTH
			}
		default:
			panic(fmt.Sprintf("Unknown dir: %c", dir))
		}

		switch facing {
		case NORTH:
			curr.y += amount
		case SOUTH:
			curr.y -= amount
		case EAST:
			curr.x += amount
		case WEST:
			curr.x -= amount
		}
	}

	return fmt.Sprintf("%d", int(math.Abs(float64(curr.x))+math.Abs(float64(curr.y))))
}

func (_d day01) b(input string) string {
	type Point struct {
		x int
		y int
	}

	const (
		NORTH = iota
		EAST
		SOUTH
		WEST
	)

	seen := make(map[Point]int)
	curr := Point{0, 0}
	facing := NORTH

	seen[curr] = 1

	for _, instr := range strings.Split(strings.Trim(input, " \n"), ", ") {
		dir := instr[0]
		amount, err := strconv.Atoi(instr[1:])
		if err != nil {
			panic(err)
		}

		switch dir {
		case 'L':
			switch facing {
			case NORTH:
				facing = WEST
			case WEST:
				facing = SOUTH
			case SOUTH:
				facing = EAST
			case EAST:
				facing = NORTH
			}
		case 'R':
			switch facing {
			case NORTH:
				facing = EAST
			case WEST:
				facing = NORTH
			case SOUTH:
				facing = WEST
			case EAST:
				facing = SOUTH
			}
		default:
			panic(fmt.Sprintf("Unknown dir: %c", dir))
		}

		for j := 0; j < amount; j++ {
			switch facing {
			case NORTH:
				curr.y += 1
			case SOUTH:
				curr.y -= 1
			case EAST:
				curr.x += 1
			case WEST:
				curr.x -= 1
			}

			seen[curr] += 1

			if seen[curr] == 2 {
				break
			}
		}

		if seen[curr] == 2 {
			break
		}
	}

	return fmt.Sprintf("%d", int(math.Abs(float64(curr.x))+math.Abs(float64(curr.y))))
}
