// Space Image Format
package sif

import (
	"image"
	"image/png"
	"os"
	"strconv"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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

	img := &image.Gray{Pix: pix, Stride: effectiveWidth, Rect: image.Rect(0, 0, effectiveWidth, effectiveHeight)}
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func (s *SIF) Visualize() {
	pixels := [][]int{}
	layerMask := [][]bool{}
	images := []*image.Gray{}

	for u := 1; u <= s.Height; u++ {
		// Allow all points at start
		for i := s.Height - u; i < s.Height; i++ {
			pixels = append(pixels, []int{})
			layerMask = append(layerMask, []bool{})
			for j := 0; j < s.Width; j++ {
				pixels[i] = append(pixels[i], 0)
				layerMask[i] = append(layerMask[i], true)
			}
		}

		for _, layer := range s.Layers[s.Height-u:] {
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

		img := &image.Gray{Pix: pix, Stride: effectiveWidth, Rect: image.Rect(0, 0, effectiveWidth, effectiveHeight)}
		images = append(images, img)
	}
	pixelgl.Run(func() {
		cfg := pixelgl.WindowConfig{
			Title:  "Pixel Rocks!",
			Bounds: pixel.R(0, 0, 1024, 768),
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		index := 0

		for !win.Closed() {
			if index >= len(images) {
				break
			}
			pic := pixel.PictureDataFromImage(images[index])
			index += 1

			sprite := pixel.NewSprite(pic, pic.Bounds())
			mat := pixel.IM
			mat = mat.Moved(win.Bounds().Center())
			mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(10, 10))
			sprite.Draw(win, mat)
			win.Update()
			time.Sleep(time.Second)
		}
	})
}
