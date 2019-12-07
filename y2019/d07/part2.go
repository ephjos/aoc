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

type ChanBundle struct {
	Input, Output, Complete chan int
}

func MakeChanBundle() *ChanBundle {
	return &ChanBundle{
		make(chan int),
		make(chan int),
		make(chan int),
	}
}

func getMax(ic *IntCode) {
	phaseSettings := []int{5, 6, 7, 8, 9}
	max := math.MinInt32

	permutations(phaseSettings, func(arr []int) {

		icA := ic
		aChan := MakeChanBundle()
		go func() {
			icA.Compute(aChan.Input, aChan.Output, aChan.Complete)
		}()

		icB := ic
		bChan := MakeChanBundle()
		go func() {
			icB.Compute(bChan.Input, bChan.Output, bChan.Complete)
		}()

		icC := ic
		cChan := MakeChanBundle()
		go func() {
			icC.Compute(cChan.Input, cChan.Output, cChan.Complete)
		}()

		icD := ic
		dChan := MakeChanBundle()
		go func() {
			icD.Compute(dChan.Input, dChan.Output, dChan.Complete)
		}()

		icE := ic
		eChan := MakeChanBundle()
		go func() {
			icE.Compute(eChan.Input, eChan.Output, eChan.Complete)
		}()

		go func() {
			aChan.Input <- arr[0]
			bChan.Input <- arr[1]
			cChan.Input <- arr[2]
			dChan.Input <- arr[3]
			eChan.Input <- arr[4]

			aChan.Input <- 0
			for {
				bChan.Input <- <-aChan.Output
				cChan.Input <- <-bChan.Output
				dChan.Input <- <-cChan.Output
				eChan.Input <- <-dChan.Output
				aChan.Input <- <-eChan.Output
			}
		}()

		final := <-eChan.Complete

		if final > max {
			max = final
		}

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
