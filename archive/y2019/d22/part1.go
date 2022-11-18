// Part 1

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
}

type Deck = []Card

func NewDeck() Deck {
	d := make(Deck, DECK_LENGTH)

	for i := range d {
		d[i].value = i
	}

	return d
}

func DeckIndexOf(d Deck, card Card) int {
	for i, c := range d {
		if c == card {
			return i
		}
	}

	return -1
}

func DealIntoNewStack(d Deck) Deck {
	out := make(Deck, DECK_LENGTH)

	l := DECK_LENGTH - 1
	for i := l; i > 0; i-- {
		out[l-i] = d[i]
	}

	return out
}

func CutN(d Deck, n int) Deck {
	out := make(Deck, 0)

	if n < 0 {
		n = DECK_LENGTH + n
	}

	head := d[:n]
	tail := d[n:]
	out = append(out, tail...)
	out = append(out, head...)

	return out
}

func DealWithIncrementN(d Deck, n int) Deck {
	out := make(Deck, DECK_LENGTH)

	count := 0
	i := 0

	for _, c := range d {
		out[i] = c

		i += n
		i %= DECK_LENGTH

		count++
	}

	return out
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	d := NewDeck()

	for input.Scan() {
		line := input.Text()
		tok := strings.Split(line, " ")

		// deal into new stack
		if tok[0] == "deal" &&
			tok[1] == "into" &&
			tok[2] == "new" &&
			tok[3] == "stack" {
			d = DealIntoNewStack(d)
		}

		// cut n
		if tok[0] == "cut" {
			n, _ := strconv.Atoi(tok[1])
			d = CutN(d, n)
		}

		// deal with increment n
		if tok[0] == "deal" &&
			tok[1] == "with" &&
			tok[2] == "increment" {
			n, _ := strconv.Atoi(tok[3])
			d = DealWithIncrementN(d, n)
		}

	}

	fmt.Println(DeckIndexOf(d, Card{2019}))
}
