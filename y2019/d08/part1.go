// Part 1

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func parseImage(data string) [][][]int {
	width := 25
	height := 6

	layers := [][][]int{}
	layer := [][]int{}
	layerPart := []int{}

	for _, stringValue := range data {
		intValue, _ := strconv.Atoi(string(stringValue))
		layerPart = append(layerPart, intValue)

		if len(layerPart) == width {
			layer = append(layer, layerPart)
			layerPart = []int{}
		}

		if len(layer) == height {
			layers = append(layers, layer)
			layer = [][]int{}
		}
	}

	return layers
}

func verifyLayer(image [][][]int) {
	min := math.MaxInt32
	minIndex := 0

	for i, layer := range image {
		zeroCount := 0
		for _, part := range layer {
			for _, value := range part {
				if value == 0 {
					zeroCount += 1
				}
			}
		}

		if zeroCount < min {
			min = zeroCount
			minIndex = i
		}
	}

	layer := image[minIndex]
	oneCount := 0
	twoCount := 0

	for _, part := range layer {
		for _, value := range part {
			switch value {
			case 1:
				oneCount += 1
			case 2:
				twoCount += 1
			default:
				break
			}
		}
	}

	fmt.Println(oneCount * twoCount)

}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()

	image := parseImage(line)
	verifyLayer(image)
}
