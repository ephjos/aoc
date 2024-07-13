package main

import (
	"math"
	"strings"
)

type day06 struct{}

func (_d day06) a(input string) string {
	// Find length of first word to enable allocation
	in := strings.Trim(input, " \n")
	l := 0
	for in[l] != '\n' {
		l += 1
	}

	// Allocate
	freqs := make([][]int, l)
	for i := range freqs {
		freqs[i] = make([]int, 26)
	}

	// Iterate over all characters, skipping new lines
	j := 0
	for _, c := range strings.Trim(input, " \n") {
		if c == '\n' {
			j = 0
			continue
		}
		freqs[j][c-97] += 1
		j += 1
	}

	// Build output
	out := strings.Builder{}
	for _, letter_freqs := range freqs {
		m := 0
		mi := 0
		for j, f := range letter_freqs {
			if f > m {
				m = f
				mi = j
			}
		}
		out.WriteRune(rune(mi + 97))
	}

	return out.String()
}

func (_d day06) b(input string) string {

	// Find length of first word to enable allocation
	in := strings.Trim(input, " \n")
	l := 0
	for in[l] != '\n' {
		l += 1
	}

	// Allocate
	freqs := make([][]int, l)
	for i := range freqs {
		freqs[i] = make([]int, 26)
	}

	// Iterate over all characters, skipping new lines
	j := 0
	for _, c := range strings.Trim(input, " \n") {
		if c == '\n' {
			j = 0
			continue
		}
		freqs[j][c-97] += 1
		j += 1
	}

	// Build output
	out := strings.Builder{}
	for _, letter_freqs := range freqs {
		m := math.MaxInt32
		mi := 0
		for j, f := range letter_freqs {
			if f < m && f != 0 {
				m = f
				mi = j
			}
		}
		out.WriteRune(rune(mi + 97))
	}

	return out.String()
}
