package mattebuilder

import (
	"image"
	"image/color"
)

func BakeAO(imgDiff, imgAO image.Image, intensity float64) image.Image {
	bounds := imgDiff.Bounds()
	out := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			dr, dg, db, da := imgDiff.At(x, y).RGBA()
			ar, _, _, _ := imgAO.At(x, y).RGBA()

			diffR := float64(dr >> 8)
			diffG := float64(dg >> 8)
			diffB := float64(db >> 8)
			aoFactor := float64(ar>>8) / 255.0

			aoFactor = aoFactor*intensity + (1.0 - intensity)

			out.SetNRGBA(x, y, color.NRGBA{
				R: uint8(diffR * aoFactor),
				G: uint8(diffG * aoFactor),
				B: uint8(diffB * aoFactor),
				A: uint8(da >> 8),
			})
		}
	}
	return out
}
