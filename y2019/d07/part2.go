// Part 1

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func permutations(arr []int, f func([]int)) {
	perm_core(arr, f, 0)
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func perm_core(arr []int, f func([]int), i int) {
	if i > len(arr) {
		f(arr)
		return
	}
	perm_core(arr, f, i+1)

	for j := i + 1; j < len(arr); j++ {
		swap(arr, i, j)
		perm_core(arr, f, i+1)
		swap(arr, i, j)
	}
}

func getMax(ic *IntCode) {
	phaseSettings := []int{5, 6, 7, 8, 9}
	max := math.MinInt32

	permutations(phaseSettings, func(arr []int) {
	})

	fmt.Println(max)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	tokens := strings.Split(line, ",")

	ic := MakeIntCode(tokens)

	getMax(ic)

}
