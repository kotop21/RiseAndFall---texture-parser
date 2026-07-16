package main

import (
	"bytes"
	"image/png"
	"os"

	"github.com/ftrvxmtrx/tga"
)

func unpackTGA(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	img, err := tga.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return err
	}

	return os.WriteFile(dst, buf.Bytes(), 0644)
}

func packTGA(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tga.Encode(&buf, img); err != nil {
		return err
	}

	return os.WriteFile(dst, buf.Bytes(), 0644)
}
