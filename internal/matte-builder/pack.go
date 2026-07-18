package mattebuilder

import (
	"image"
	"image/color"
)

func PackNormalSpec(imgNormal, imgSpec image.Image) image.Image {
	bounds := imgNormal.Bounds()
	out := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			nr, ng, nb, _ := imgNormal.At(x, y).RGBA()
			sr, _, _, _ := imgSpec.At(x, y).RGBA()

			rByte := uint8(nr >> 8)
			gByte := uint8(ng >> 8)
			bByte := uint8(nb >> 8)
			aByte := uint8(sr >> 8)

			gByte = 255 - gByte

			out.SetNRGBA(x, y, color.NRGBA{
				R: rByte,
				G: gByte,
				B: bByte,
				A: aByte,
			})
		}
	}
	return out
}
