// Part 1

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Point struct {
	X float64 // Distance from left
	Y float64 // Distance from top
}

func getAsteroidPoints(scanner *bufio.Scanner) ([]Point, int, int) {
	points := make([]Point, 0)
	h := 0
	count := 0

	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, v := range line {
			if string(v) == "#" {
				points = append(points, Point{float64(i), float64(j)})
			}
			count += 1
		}
		j += 1
	}

	h = j
	w := count / j

	return points, w, h
}

func visualize(points *[]Point, w, h int) {
	pixelgl.Run(func() {
		cfg := pixelgl.WindowConfig{
			Title:  "SIF Visualizer",
			Bounds: pixel.R(0, 0, 768, 768),
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		offset := 20.
		bounds := win.Bounds()
		height := bounds.H() - (2 * offset)
		width := bounds.W() - offset

		wStep := width / float64(w)
		hStep := height / float64(h)
		offsetVec := pixel.V(offset, offset)

		imd := imdraw.New(nil)
		lines := imdraw.New(nil)

		vecs := make([]pixel.Vec, 0)

		for _, point := range *points {

			vec := pixel.V(point.X*wStep, height-(point.Y*hStep)).Add(offsetVec)
			vecs = append(vecs, vec)
			imd.Push(vec)
			imd.Circle(5, 1)
		}

		i := 0
		for !win.Closed() {
			win.Clear(colornames.Black)

			imd.Draw(win)

			if i < len(vecs) {
				lines.Clear()

				scale := 1000.
				res := 360.
				base := pixel.V(0, 0)

				for j := 0.; j < res; j++ {
					rads := j * (360. / res) * (math.Pi / 180.)

					xAng := math.Cos(rads)
					yAng := math.Sin(rads)
					x := scale * xAng
					y := scale * yAng

					base = pixel.V(x, y).Add(vecs[i])

					lines.Push(vecs[i], base)
					lines.Line(1)
				}

				lines.Draw(win)

				i += 1
				time.Sleep(time.Second / 20.)
			}

			win.Update()

		}
	})

}

func main() {
	input := bufio.NewScanner(os.Stdin)
	points, w, h := getAsteroidPoints(input)
	fmt.Println(points)
	visualize(&points, w, h)
}
