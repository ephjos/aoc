// Part 2

package main

import (
	"bufio"
	"os"

	"github.com/josephthomashines/aoc/y2019/sif"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()

	s := &sif.SIF{25, 6, nil}
	s.Parse(line)
	s.SaveImage("./testImage.png")
}
