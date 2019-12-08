package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/josephthomashines/aoc/y2019/sif"
)

func main() {
	fmt.Println("AA")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()

	fmt.Println("AA")
	s := &sif.SIF{25, 6, nil}
	s.Parse(line)
	s.Visualize()

}
