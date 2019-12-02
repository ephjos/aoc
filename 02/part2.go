// Part 1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func runIntCode(memory []int, noun, verb int) int {
	memory[1] = noun
	memory[2] = verb

	for i := 0; i < len(memory); i += 4 {
		opcode := memory[i]
		if opcode == 99 {
			break
		}

		i1 := memory[i+1]
		i2 := memory[i+2]
		d := memory[i+3]

		if opcode == 1 {
			memory[d] = memory[i1] + memory[i2]
		} else if opcode == 2 {
			memory[d] = memory[i1] * memory[i2]
		}
	}

	return memory[0]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	temp := strings.Split(input.Text(), ",")
	var data []int

	for _, v := range temp {
		t, _ := strconv.Atoi(v)
		data = append(data, t)
	}

	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			var dataCopy = make([]int, len(data))
			copy(dataCopy, data)
			output := runIntCode(dataCopy, i, j)
			if output == 19690720 {
				fmt.Println(100*i + j)
				return
			}
		}
	}
}
