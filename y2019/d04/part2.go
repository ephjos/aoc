// Part 1

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ascending(digits []string) bool {
	lastVal := -1
	for _, digit := range digits {
		val, _ := strconv.Atoi(digit)
		if val < lastVal {
			return false
		}
		lastVal = val
	}
	return true
}

func onePair(digits []string) bool {
	b3 := -3
	b2 := -2
	b1 := -1

	digits = append(digits, strconv.Itoa(math.MaxInt32))

	for _, digit := range digits {
		val, _ := strconv.Atoi(digit)
		if val != b1 && b1 == b2 && b2 != b3 {
			return true
		}
		b3 = b2
		b2 = b1
		b1 = val
	}
	return false
}

func getPotential(low, high int) {
	count := 0
	rules := []func([]string) bool{
		ascending,
		onePair,
	}

	for i := low; i < high; i++ {
		digits := strings.Split(strconv.Itoa(i), "")
		allRulesFollowed := true
		for _, rule := range rules {
			if !rule(digits) {
				allRulesFollowed = false
				break
			}
		}
		if allRulesFollowed {
			count += 1
		}
	}

	fmt.Println(count)
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	input.Scan()
	line := input.Text()
	numbers := strings.Split(line, "-")
	low, _ := strconv.Atoi(numbers[0])
	high, _ := strconv.Atoi(numbers[1])

	getPotential(low, high)
}
