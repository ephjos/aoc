// Part 2

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
	RUN  = "RUN"
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

	// Write final RUN command
	for _, r := range RUN {
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
		"NOT H J",
		"OR C J",
		"AND B J",
		"AND A J",
		"NOT J J",
		"AND D J",
	}

	runComputer(ic, inputs)
}
