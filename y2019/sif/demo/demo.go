package main

import (
	"bufio"
	"os"

	"github.com/ephjos/aoc/y2019/sif"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()

	s := &sif.SIF{25, 6, nil}
	s.Parse(line)
	s.Visualize()
}
