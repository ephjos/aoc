// Part 2

package main

import (
	"bufio"
	"image"
	"image/png"
	"os"
	"strconv"
)

const (
	WIDTH  int = 25
	HEIGHT int = 6
)

func parseImage(data string) [][][]int {

	layers := [][][]int{}
	layer := [][]int{}
	layerPart := []int{}

	for _, stringValue := range data {
		intValue, _ := strconv.Atoi(string(stringValue))
		layerPart = append(layerPart, intValue)

		if len(layerPart) == WIDTH {
			layer = append(layer, layerPart)
			layerPart = []int{}
		}

		if len(layer) == HEIGHT {
			layers = append(layers, layer)
			layer = [][]int{}
		}
	}

	return layers
}

func collapseLayers(imageData [][][]int) {
	pixels := [][]int{}
	layerMask := [][]bool{}

	// Allow all points at start
	for i := 0; i < HEIGHT; i++ {
		pixels = append(pixels, []int{})
		layerMask = append(layerMask, []bool{})
		for j := 0; j < WIDTH; j++ {
			pixels[i] = append(pixels[i], 0)
			layerMask[i] = append(layerMask[i], true)
		}
	}

	for _, layer := range imageData {
		for i, part := range layer {
			for j, v := range part {
				if layerMask[i][j] {
					switch v {
					case 0:
						pixels[i][j] = 0
						layerMask[i][j] = false
					case 1:
						pixels[i][j] = 255
						layerMask[i][j] = false
					default:
						break
					}
				}
			}
		}
	}

	pix := []byte{}
	for _, row := range pixels {
		for _, value := range row {
			pix = append(pix, byte(value))
		}
	}

	img := &image.Gray{Pix: pix, Stride: WIDTH, Rect: image.Rect(0, 0, WIDTH, HEIGHT)}
	f, err := os.Create("./image")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)

}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()

	image := parseImage(line)
	collapseLayers(image)
}
