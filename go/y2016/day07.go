package main

import (
	"fmt"
	"strings"
)

type day07 struct{}

func (_d day07) a(input string) string {
	count := 0
	for _, line := range strings.Split(strings.Trim(input, " \n"), "\n") {
		abba_out := false
		abba_in := false

		in_hypernet := false
		for i := 0; i < len(line)-3; i++ {
			a := line[i]
			b := line[i+1]
			c := line[i+2]
			d := line[i+3]

			if a == '[' {
				in_hypernet = true
				continue
			} else if a == ']' {
				in_hypernet = false
				continue
			}

			if a != b && c != d && a == d && b == c {
				if in_hypernet {
					abba_in = true
				} else {
					abba_out = true
				}
			}
		}

		if abba_out && !abba_in {
			count += 1
		}
	}

	return fmt.Sprint(count)
}

func (_d day07) b(input string) string {
	count := 0
outer:
	for _, line := range strings.Split(strings.Trim(input, " \n"), "\n") {
		x := [26][26]uint8{}

		in_hypernet := false
		for i := 0; i < len(line)-2; i++ {
			a := line[i]
			b := line[i+1]
			c := line[i+2]

			if a == '[' {
				in_hypernet = true
			} else if a == ']' {
				in_hypernet = false
			}

			if a == '[' || a == ']' || b == '[' || b == ']' || c == '[' || c == ']' {
				continue
			}

			if a != b && b != c && a == c {
				if in_hypernet {
					x[a-97][b-97] |= 0x0F
				} else {
					x[b-97][a-97] |= 0xF0
				}
			}
		}

		for i := range 26 {
			for j := range 26 {
				if x[i][j] == 255 {
					count += 1
					continue outer
				}
			}
		}

	}

	return fmt.Sprint(count)
}
