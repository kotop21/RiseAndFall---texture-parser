package mattefixer

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func ModifyAlpha(pngPath string, strength int) error {
	file, err := os.Open(pngPath)
	if err != nil {
		return err
	}

	img, err := png.Decode(file)
	file.Close()
	if err != nil {
		return err
	}

	bounds := img.Bounds()
	outImg := image.NewRGBA(bounds)

	alphaVal := uint8((100 - strength) * 255 / 100)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			r, g, b, _ := originalColor.RGBA()
			outImg.Set(x, y, color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: alphaVal,
			})
		}
	}

	outFile, err := os.Create(pngPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, outImg)
}
