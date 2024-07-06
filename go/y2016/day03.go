package main

import (
	"fmt"
	"strconv"
	"strings"
)

type day03 struct{}

func (_d day03) a(input string) string {
	count := 0

	for _, line := range strings.Split(strings.Trim(input, " \n"), "\n") {
		toks := strings.Fields(line)
		x, _ := strconv.Atoi(toks[0])
		y, _ := strconv.Atoi(toks[1])
		z, _ := strconv.Atoi(toks[2])

		if x+y > z && y+z > x && x+z > y {
			count += 1
		}
	}

	return fmt.Sprint(count)
}

func (_d day03) b(input string) string {
	count := 0
	lines := strings.Split(strings.Trim(input, " \n"), "\n")
	n := len(lines)

	c := make([]int, n*3)

	for i, line := range lines {
		toks := strings.Fields(line)
		x, _ := strconv.Atoi(toks[0])
		y, _ := strconv.Atoi(toks[1])
		z, _ := strconv.Atoi(toks[2])

		c[i] = x
		c[n+i] = y
		c[(2*n)+i] = z
	}

	for i := 0; i < len(c); i += 3 {
		x := c[i]
		y := c[i+1]
		z := c[i+2]

		if x+y > z && y+z > x && x+z > y {
			count += 1
		}
	}

	return fmt.Sprint(count)
}
