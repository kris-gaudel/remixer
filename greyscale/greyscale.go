package greyscale

import (
	"image"
	"image/color"
)

type GreyscaleFilter struct {
	image.Image
}

func (f *GreyscaleFilter) At(x, y int) color.Color {
	r, g, b, a := f.Image.At(x, y).RGBA()
	grey := uint16(float64(r)*0.21 + float64(g)*0.72 + float64(b)*0.07)

	return color.RGBA64 {
		R: grey,
		G: grey,
		B: grey,
		A: uint16(a),
	}
}