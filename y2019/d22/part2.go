// Part 2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	DECK_LENGTH = 10007
)

type Card struct {
	value int
	index int
}

func DealIntoNewStack(c *Card) {
	c.index = DECK_LENGTH - 1 - c.index
}

func CutN(c *Card, n int) {
	if c.index < n {
		c.index += DECK_LENGTH - n
	} else {
		c.index -= n
	}
}

func DealWithIncrementN(c *Card, n int) {
	i := 0

	for idx := 0; idx < DECK_LENGTH; idx++ {
		if c.index == idx {
			c.index = i
			return
		}

		i += n
		i %= DECK_LENGTH
	}

}

func main() {
	input := bufio.NewScanner(os.Stdin)

	val := 2019
	c := Card{val, val}

	for input.Scan() {
		line := input.Text()
		tok := strings.Split(line, " ")

		// deal into new stack
		if tok[0] == "deal" &&
			tok[1] == "into" &&
			tok[2] == "new" &&
			tok[3] == "stack" {
			DealIntoNewStack(&c)
		}

		// cut n
		if tok[0] == "cut" {
			n, _ := strconv.Atoi(tok[1])
			CutN(&c, n)
		}

		// deal with increment n
		if tok[0] == "deal" &&
			tok[1] == "with" &&
			tok[2] == "increment" {
			n, _ := strconv.Atoi(tok[3])
			DealWithIncrementN(&c, n)
		}

		fmt.Println(c)
	}

	fmt.Println(c)
}
