package mattebuilder

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func OpenImage(path string) (image.Image, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(path)
	var img image.Image
	reader := bytes.NewReader(data)

	if ext == ".jpg" || ext == ".jpeg" {
		img, err = jpeg.Decode(reader)
	} else {
		img, err = png.Decode(reader)
	}
	if err != nil {
		return nil, err
	}

	return img, nil
}
