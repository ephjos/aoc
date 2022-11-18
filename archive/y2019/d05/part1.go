// Part 1

package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	intCode := MakeIntCode(strings.Split(input.Text(), ","))
	intCode.Compute(1)
}
