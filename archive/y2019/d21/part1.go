// Part 1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	NL   = 10
	WALK = "WALK"
)

func runComputer(ic *IntCode, inputs []string) {
	// Add inputs
	for _, inp := range inputs {
		runes := []rune(inp)

		for _, r := range runes {
			ic.AddInput(int(r))
		}

		ic.AddInput(NL)
	}

	// Write final WALK command
	for _, r := range WALK {
		ic.AddInput(int(r))
	}
	ic.AddInput(NL)

	// Run computer
	for ic.IsRunning {
		out := ic.Run()

		if out < 130 { // Character
			fmt.Print(string(out))
		} else { // Number
			fmt.Println(out)
		}
	}

}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	tokens := strings.Split(line, ",")

	ic := NewIntCode(tokens)

	inputs := []string{
		"NOT A T",
		"NOT B J",
		"OR J T",
		"NOT C J",
		"OR J T",
		"NOT D J",
		"NOT J J",
		"AND T J",
	}

	runComputer(ic, inputs)
}
