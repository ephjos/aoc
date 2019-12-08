# SIF Demo

I made a miny library for the Space Image Format (SIF)
from Day 8. This `demo.go` file will output a visualization
of the data stream when ran like:
	```bash
	go run demo.go < input
	```

This starts from the bottom of the layers and builds the
image from the bottom up. I used `pixelgl` to render
each image to the screen.
