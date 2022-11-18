// Part 2

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

	ic := New(tokens)

	ch := make(chan int64)

	go ic.Compute(ch)

	ch <- 2

	for x := range ch {
		fmt.Println(x)
	}

}
