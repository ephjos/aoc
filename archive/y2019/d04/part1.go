// Part 1

package main

import (
	"bufio"
	"fmt"
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

func atLeastOnePair(digits []string) bool {
	lastVal := -1
	for _, digit := range digits {
		val, _ := strconv.Atoi(digit)
		if val == lastVal {
			return true
		}
		lastVal = val
	}
	return false
}

func getPotential(low, high int) {
	count := 0
	rules := []func([]string) bool{
		ascending,
		atLeastOnePair,
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
