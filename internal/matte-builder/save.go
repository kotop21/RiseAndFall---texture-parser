package mattebuilder

import (
	"image"
	"image/png"
	"os"
)

func SavePNG(path string, img image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}
