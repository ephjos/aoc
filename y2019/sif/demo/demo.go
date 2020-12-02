package main

import (
	"bufio"
	"os"

	sif ".."
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()

	s := &sif.SIF{25, 6, nil}
	s.Parse(line)
	s.Visualize()
}
