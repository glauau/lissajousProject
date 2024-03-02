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
	color.RGBA{R: 0x24, G: 0x6A, B: 0x73, A: 0xFF}, // #246A73
	color.RGBA{R: 0x36, G: 0x8F, B: 0x8B, A: 0xFF}, // #368F8B
	color.RGBA{R: 0xF3, G: 0xDF, B: 0xC1, A: 0xFF}, // #F3DFC1
}

const (
	blackIndex = 0
	blueIndex  = 1
	greenIndex = 2
	beigeIndex = 3

	cycles  = 6
	res     = 0.001
	size    = 300
	nframes = 130
	delay   = 30
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
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
			// Alternating between different colors based on the frame number
			switch i % 3 {
			case 0:
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blueIndex+1)
			case 1:
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex+1)
			case 2:
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), beigeIndex+1)
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
