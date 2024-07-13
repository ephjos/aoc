package main

/*
import (
	"math"
	"strconv"
	"strings"
)
*/

import (
	"fmt"
	"strconv"
	"strings"
)

type day08 struct{}

func (_d day08) a(input string) string {
	const screen_h = 6
	const screen_w = 50
	screen := [screen_h][screen_w]uint8{}

	for _, line := range strings.Split(strings.Trim(input, " \n"), "\n") {
		toks := strings.Split(line, " ")
		switch toks[1] {
		case "column":
			n_toks := strings.Split(toks[2], "=")
			col, _ := strconv.Atoi(n_toks[1])
			shift, _ := strconv.Atoi(toks[4])

			new_col := [screen_h]uint8{}

			for i := range screen_h {
				new_col[(i+shift)%screen_h] = screen[i][col]
			}

			for i := range screen_h {
				screen[i][col] = new_col[i]
			}
			break
		case "row":
			n_toks := strings.Split(toks[2], "=")
			row, _ := strconv.Atoi(n_toks[1])
			shift, _ := strconv.Atoi(toks[4])

			new_row := [screen_w]uint8{}

			for j := range screen_w {
				new_row[(j+shift)%screen_w] = screen[row][j]
			}

			for j := range screen_w {
				screen[row][j] = new_row[j]
			}
			break
		default:
			// rect
			n_toks := strings.Split(toks[1], "x")
			cols, _ := strconv.Atoi(n_toks[0])
			rows, _ := strconv.Atoi(n_toks[1])

			for i := range rows {
				for j := range cols {
					screen[i][j] = 1
				}
			}

			break
		}
	}

	count := 0

	for i := range screen_h {
		for j := range screen_w {
			count += int(screen[i][j])
		}
	}

	return fmt.Sprint(count)
}

func (_d day08) b(input string) string {
	const screen_h = 6
	const screen_w = 50
	screen := [screen_h][screen_w]uint8{}

	for _, line := range strings.Split(strings.Trim(input, " \n"), "\n") {
		toks := strings.Split(line, " ")
		switch toks[1] {
		case "column":
			n_toks := strings.Split(toks[2], "=")
			col, _ := strconv.Atoi(n_toks[1])
			shift, _ := strconv.Atoi(toks[4])

			new_col := [screen_h]uint8{}

			for i := range screen_h {
				new_col[(i+shift)%screen_h] = screen[i][col]
			}

			for i := range screen_h {
				screen[i][col] = new_col[i]
			}
			break
		case "row":
			n_toks := strings.Split(toks[2], "=")
			row, _ := strconv.Atoi(n_toks[1])
			shift, _ := strconv.Atoi(toks[4])

			new_row := [screen_w]uint8{}

			for j := range screen_w {
				new_row[(j+shift)%screen_w] = screen[row][j]
			}

			for j := range screen_w {
				screen[row][j] = new_row[j]
			}
			break
		default:
			// rect
			n_toks := strings.Split(toks[1], "x")
			cols, _ := strconv.Atoi(n_toks[0])
			rows, _ := strconv.Atoi(n_toks[1])

			for i := range rows {
				for j := range cols {
					screen[i][j] = 1
				}
			}

			break
		}
	}

	c1 := "█"
	c2 := " "

	// Display the screen
	for _ = range screen_w + (screen_w / 5) + 1 {
		fmt.Print(c2)
	}
	fmt.Println()

	for i := range screen_h {
		fmt.Print(c2)
		for j := range screen_w {
			if screen[i][j] == 1 {
				fmt.Print(c1)
			} else {
				fmt.Print(c2)
			}

			if (j+1)%5 == 0 {
				fmt.Print(c2)
			}
		}
		fmt.Println()
	}
	for _ = range screen_w + (screen_w / 5) + 1 {
		fmt.Print(c2)
	}
	fmt.Println()

	return "^"
}
