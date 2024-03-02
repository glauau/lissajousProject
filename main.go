package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}, // Black
	color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}, // Red
	color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}, // Green
	color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xFF}, // Blue
}

const (
	blackIndex = 0
	redIndex   = 1
	greenIndex = 2
	blueIndex  = 3
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 4
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// Alternate between different colors based on the position of x and y
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(i%len(palette))+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
