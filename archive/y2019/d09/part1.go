// Part 1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	tokens := strings.Split(line, ",")

	ic := MakeIntCode(tokens)

	ch := make(chan int)

	go ic.Compute(ch)

	ch <- 1

	for x := range ch {
		fmt.Println(x)
	}

}
