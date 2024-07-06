package main

import (
	"fmt"
	"strconv"
	"strings"
)

type day04 struct{}

func (_d day04) a(input string) string {
	count := 0

	for _, line := range strings.Split(strings.Trim(input, " \n"), "\n") {
		counts := [26]int{}
		i := 0
		sector_buf := strings.Builder{}

		for j, d := range line {
			if d == '-' {
				continue
			}

			if d >= '0' && d <= '9' {
				sector_buf.WriteRune(rune(d))
				continue
			}

			if d == '[' {
				i = j + 1
				break
			}

			counts[d-97] += 1
		}

		sector, _ := strconv.Atoi(sector_buf.String())

		is_real := true
		for j := range 5 {
			m := 0
			mi := 0

			for k, c := range counts {
				if c > m {
					m = c
					mi = k
				}
			}

			counts[mi] = 0

			if byte(mi+97) != line[i+j] {
				is_real = false
				break
			}
		}

		if is_real {
			count += sector
		}
	}

	return fmt.Sprint(count)
}

func (_d day04) b(input string) string {

	for _, line := range strings.Split(strings.Trim(input, " \n"), "\n") {
		counts := [26]int{}
		i := 0
		sector_buf := strings.Builder{}

		for j, d := range line {
			if d == '-' {
				continue
			}

			if d >= '0' && d <= '9' {
				sector_buf.WriteRune(rune(d))
				continue
			}

			if d == '[' {
				i = j + 1
				break
			}

			counts[d-97] += 1
		}

		sector, _ := strconv.Atoi(sector_buf.String())

		is_real := true
		for j := range 5 {
			m := 0
			mi := 0

			for k, c := range counts {
				if c > m {
					m = c
					mi = k
				}
			}

			counts[mi] = 0

			if byte(mi+97) != line[i+j] {
				is_real = false
				break
			}
		}

		if is_real {
			room_buf := strings.Builder{}
			for _, c := range line {
				if c == '-' {
					room_buf.WriteRune(' ')
					continue
				}

				if c == '[' || (c >= '0' && c <= '9') {
					break
				}

				room_buf.WriteRune(rune(((int(c-97) + sector) % 26) + 97))
			}

			room := strings.Trim(room_buf.String(), " \n")

			if room == "northpole object storage" {
				return fmt.Sprint(sector)
			}
		}
	}

	return "---"
}
