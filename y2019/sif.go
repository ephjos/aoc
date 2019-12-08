// Space Image Format
package sif

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"strconv"
)

type SIF struct {
	Width  int
	Height int
	Layers [][][]int
}

// Parse takes the raw data "stream" and makes it
// into a "layers" array
func (s *SIF) Parse(data string) {
	layers := [][][]int{}
	layer := [][]int{}
	layerPart := []int{}

	for _, stringValue := range data {
		intValue, _ := strconv.Atoi(string(stringValue))
		layerPart = append(layerPart, intValue)

		if len(layerPart) == s.Width {
			layer = append(layer, layerPart)
			layerPart = []int{}
		}

		if len(layer) == s.Height {
			layers = append(layers, layer)
			layer = [][]int{}
		}
	}

	s.Layers = layers
}

func (s *SIF) SaveImage(filename string) {
	pixels := [][]int{}
	layerMask := [][]bool{}

	// Allow all points at start
	for i := 0; i < s.Height; i++ {
		pixels = append(pixels, []int{})
		layerMask = append(layerMask, []bool{})
		for j := 0; j < s.Width; j++ {
			pixels[i] = append(pixels[i], 0)
			layerMask[i] = append(layerMask[i], true)
		}
	}

	for _, layer := range s.Layers {
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
	padding := 10
	effectiveWidth := s.Width + (2 * padding)
	effectiveHeight := s.Height + padding
	for i := 0; i < padding/2; i++ {
		for j := 0; j < effectiveWidth; j++ {
			pix = append(pix, 0)
		}
	}

	for _, row := range pixels {
		for i := 0; i < padding; i++ {
			pix = append(pix, 0)
		}
		for _, value := range row {
			pix = append(pix, byte(value))
		}
		for i := 0; i < padding; i++ {
			pix = append(pix, 0)
		}
	}

	for i := 0; i < padding/2; i++ {
		for j := 0; j < effectiveWidth; j++ {
			pix = append(pix, 0)
		}
	}

	fmt.Println(pix)
	img := &image.Gray{Pix: pix, Stride: effectiveWidth, Rect: image.Rect(0, 0, effectiveWidth, effectiveHeight)}
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
