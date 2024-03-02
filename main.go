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

var palette = []color.Color{color.Black, color.RGBA{R: 0xFF, G: 0x68, B: 0x6B, A: 0xFF}, color.RGBA{R: 0xA5, G: 0xFF, B: 0xD6, A: 0xFF}}

const (
	blackIndex = 0
	redIndex   = 1
	greenIndex = 2
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
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
		// Fill the background with black color
		for y := 0; y < 2*size+1; y++ {
			for x := 0; x < 2*size+1; x++ {
				img.SetColorIndex(x, y, blackIndex)
			}
		}
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// Draw red and green lines using the specified colors
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), redIndex)
			img.SetColorIndex(size+int(y*size+0.5), size+int(x*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
