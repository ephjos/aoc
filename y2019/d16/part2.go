// Part 2

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func makePattern(position, length int) []int {
	basePattern := []int{0, 1, 0, -1}

	output := make([]int, 0, length)

outerLoop:
	for {
		for _, d := range basePattern {
			for i := 0; i < position+1; i++ {
				output = append(output, d)
				if len(output) > length {
					break outerLoop
				}
			}
		}
	}

	return output[1:]
}

func dot(s1, s2 []int) int {
	if len(s1) != len(s2) {
		panic("Lengths mismatched...")
	}

	sum := 0

	for i := range s1 {
		sum += s1[i] * s2[i]
	}

	o := math.Abs(float64(sum % 10.))

	return int(o)
}

func applyPattern(list []int) []int {
	output := make([]int, 0)

	length := len(list)
	for i := range list {
		pattern := makePattern(i, length)
		newDigit := dot(pattern, list)
		output = append(output, newDigit)
	}

	return output
}

func sliceToInt(slice []int) int {
	length := len(slice)

	sum := 0
	count := 0.
	for i := length - 1; i >= 0; i-- {
		exp := int(math.Pow(10, count))

		sum += slice[i] * exp

		count += 1.
	}

	return sum
}

func iterate(n, offset int, list []int) int {
	output := make([]int, 0)
	output = list[offset:]

	for i := 0; i < n; i++ {
		sum := 0
		for j := len(output) - 1; j >= 0; j-- {
			sum += output[j]
			output[j] = sum % 10
		}

	}

	output = output[:8]

	return sliceToInt(output)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	tokens := strings.Split(line, "")

	list := make([]int, len(tokens))
	for i, s := range tokens {
		list[i], _ = strconv.Atoi(s)
	}

	fullList := make([]int, 0)
	for i := 0; i < 10000; i++ {
		fullList = append(fullList, list...)
	}

	offset := sliceToInt(fullList[:7])
	fmt.Println(iterate(100, offset, fullList))

}
