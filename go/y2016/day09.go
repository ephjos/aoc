package main

import (
	"fmt"
	"strings"
)

import ()

type day09 struct{}

func (_d day09) a(input string) string {
	compressed := strings.Trim(input, " \n")
	count := 0

	i := 0
	for i < len(compressed) {
		c := compressed[i]

		if c == '(' {
			end_paren := i
			for compressed[end_paren] != ')' {
				end_paren += 1
			}

			l := 0
			j := i + 1
			for compressed[j] != 'x' {
				l *= 10
				l += int(compressed[j] - 48)
				j++
			}

			j++
			r := 0
			for ; j < end_paren; j++ {
				r *= 10
				r += int(compressed[j] - 48)
			}

			count += l * r
			i = end_paren + l + 1
		} else {
			count += 1
			i++
		}
	}

	return fmt.Sprint(count)
}

// first cut, recursive solution
func _b__core(s string) int {
	count := 0

	i := 0
	for i < len(s) {
		c := s[i]

		if c == '(' {
			end_paren := i
			for s[end_paren] != ')' {
				end_paren += 1
			}

			l := 0
			j := i + 1
			for s[j] != 'x' {
				l *= 10
				l += int(s[j] - 48)
				j++
			}

			j++
			r := 0
			for ; j < end_paren; j++ {
				r *= 10
				r += int(s[j] - 48)
			}

			count += r * b__core(s[end_paren+1:end_paren+l+1])
			count += b__core(s[end_paren+l+1:])
			break
		} else {
			count += 1
			i += 1
		}
	}

	return count
}

type block struct {
	start int
	end   int // exclusive
	x     int
}

// manual stack management, no recursion.
// benchmarked to same perf as recursive solution, but maintain ownership
// of control flow.
func b__core(s string) int {
	// Use an array since heap allocating a slice dominates the runtime.
	// No bounds checks, will crash if overflow.
	stack := [64]block{}
	sp := 0

	stack[sp] = block{0, len(s), 1}
	sp++

	count := 0

	for sp > 0 {
		sp--
		curr := stack[sp]

		for i := curr.start; i < curr.end; i++ {
			c := s[i]

			if c == '(' {
				end_paren := i
				for s[end_paren] != ')' {
					end_paren += 1
				}

				l := 0
				j := i + 1
				for s[j] != 'x' {
					l *= 10
					l += int(s[j] - 48)
					j++
				}

				j++
				r := 0
				for ; j < end_paren; j++ {
					r *= 10
					r += int(s[j] - 48)
				}

				stack[sp].start = end_paren + 1
				stack[sp].end = end_paren + l + 1
				stack[sp].x = curr.x * r
				sp++

				stack[sp].start = end_paren + l + 1
				stack[sp].end = curr.end
				stack[sp].x = curr.x
				sp++
				break
			} else {
				count += curr.x
			}
		}
	}

	return count
}

// start from back
// track length from current index through end of string
// no recursion
func (_d day09) b(input string) string {
	compressed := strings.Trim(input, " \n")
	return fmt.Sprint(b__core(compressed))
}
