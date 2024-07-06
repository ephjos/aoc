package main

import (
	"strings"
)

type day02 struct{}

func (_d day02) a(input string) string {
	table := [26][10]rune{}
	table['U'-65] = [10]rune{'1', '2', '3', '1', '2', '3', '4', '5', '6'}
	table['D'-65] = [10]rune{'4', '5', '6', '7', '8', '9', '7', '8', '9'}
	table['L'-65] = [10]rune{'1', '1', '2', '4', '4', '5', '7', '7', '8'}
	table['R'-65] = [10]rune{'2', '3', '3', '5', '6', '6', '8', '9', '9'}

	digits := strings.Builder{}

	curr := '5'
	for _, line := range strings.Split(strings.Trim(input, " \n"), "\n") {
		for _, d := range line {
			curr = table[d-65][curr-49]
		}

		digits.WriteRune(curr)
	}

	return digits.String()
}

func (_d day02) b(input string) string {
	table := [26][20]rune{}
	table['U'-65] = [20]rune{'1', '2', '1', '4', '5', '2', '3', '4', '9', '0', '0', '0', '0', '0', '0', '0', '6', '7', '8', 'B'}
	table['D'-65] = [20]rune{'3', '6', '7', '8', '5', 'A', 'B', 'C', '9', '0', '0', '0', '0', '0', '0', '0', 'A', 'D', 'C', 'D'}
	table['L'-65] = [20]rune{'1', '2', '2', '3', '5', '5', '6', '7', '8', '0', '0', '0', '0', '0', '0', '0', 'A', 'A', 'B', 'D'}
	table['R'-65] = [20]rune{'1', '3', '4', '4', '6', '7', '8', '9', '9', '0', '0', '0', '0', '0', '0', '0', 'B', 'C', 'C', 'D'}

	digits := strings.Builder{}

	curr := '5'
	for _, line := range strings.Split(strings.Trim(input, " \n"), "\n") {
		for _, d := range line {
			curr = table[d-65][curr-49]
		}

		digits.WriteRune(curr)
	}

	return digits.String()
}
