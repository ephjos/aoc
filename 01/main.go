package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func getFuel(moduleWeight int) int {
	out := int(math.Floor(float64(moduleWeight) / 3.))
	out -= 2

	if out <= 0 {
		return 0
	}

	return out + getFuel(out)
}

func main() {
	data, err := ioutil.ReadFile("./input")

	if err != nil {
		panic(err)
	}

	masses := strings.Split(string(data), "\n")
	masses = masses[:len(masses)-1]
	sum := 0

	for _, mass := range masses {
		intMass, err := strconv.Atoi(mass)

		if err != nil {
			panic(err)
		}

		sum += getFuel(intMass)
	}

	fmt.Println(sum)
}
