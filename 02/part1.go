// Part 1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	temp := strings.Split(input.Text(), ",")
	var data []int

	for _, v := range temp {
		t, _ := strconv.Atoi(v)
		data = append(data, t)
	}

	for i := 0; i < len(data); i += 4 {
		opcode := data[i]
		if opcode == 99 {
			break
		}

		i1 := data[i+1]
		i2 := data[i+2]
		d := data[i+3]

		if opcode == 1 {
			data[d] = data[i1] + data[i2]
		} else if opcode == 2 {
			data[d] = data[i1] * data[i2]
		}
	}

	fmt.Println(data)
}
