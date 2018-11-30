package ripper

import (
	"image"
	"image/color"
)

type Size struct {
	X int
	Y int
}

type Ripper struct {
	BlankColor *color.RGBA
	BlockSize  *Size
}

func (r *Ripper) Do(img image.Image) []image.Image {
	if r.BlankColor == nil {
		r.BlankColor = &color.RGBA{
			0x00,
			0x00,
			0x00,
			0xff,
		}
	}
	if r.BlockSize == nil {
		r.BlockSize = &Size{
			8,
			8,
		}
	}

	b := img.Bounds()

	xd := b.Max.X / r.BlockSize.X
	if 0 != b.Max.X%r.BlockSize.X {
		xd++
	}

	yd := b.Max.Y / r.BlockSize.Y
	if 0 != b.Max.Y%r.BlockSize.Y {
		yd++
	}

	ret := make([]image.Image, 0, xd*yd)

	for y := 0; y < b.Max.Y; y += r.BlockSize.Y {
		for x := 0; x < b.Max.X; x += r.BlockSize.X {
			rgba := image.NewRGBA(image.Rect(0, 0, r.BlockSize.X, r.BlockSize.Y))

			for yb := 0; yb < r.BlockSize.Y; yb++ {
				for xb := 0; xb < r.BlockSize.X; xb++ {
					if b.Max.X <= x+xb || b.Max.Y <= y+yb {
						//rgba.Set(xb, yb, r.BlankColor)
					} else {
						rgba.Set(xb, yb, img.At(x+xb, y+yb))
					}
				}
			}
			ret = append(ret, rgba)
		}
	}

	return ret
}
