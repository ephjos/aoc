package main

import (
	"fmt"
	"os"
)

type Day interface {
	a(input string) string
	b(input string) string
}

var DAYS = [...]Day{
	day01{},
	day02{},
	day03{},
	day04{},
	day05{},
}

func get_inputs() [25]string {
	var inputs [25]string

	for i := 0; i < 25; i++ {
		path := fmt.Sprintf("./inputs/d%02d", i+1)

		content, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		inputs[i] = string(content)
	}

	return inputs
}

func main() {
	// Read all inputs
	inputs := get_inputs()

	// Run DAYS
	for i, day := range DAYS {
		fmt.Printf("day%02d a: %s\n", i+1, day.a(inputs[i]))
		fmt.Printf("day%02d b: %s\n", i+1, day.b(inputs[i]))
	}
}
