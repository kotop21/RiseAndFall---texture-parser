package mattebuilder

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func OpenImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ext := filepath.Ext(path)
	var img image.Image
	if ext == ".jpg" || ext == ".jpeg" {
		img, err = jpeg.Decode(file)
	} else {
		img, err = png.Decode(file)
	}
	if err != nil {
		return nil, err
	}
	return img, nil
}
