package main

import (
	"bytes"
	"image/png"
	"os"

	"github.com/woozymasta/bcn"
)

func unpackDDS(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	_, img, err := bcn.DecodeDDS(bytes.NewReader(data))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return err
	}

	return os.WriteFile(dst, buf.Bytes(), 0644)
}

func packDDS(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}

	tex, err := bcn.EncodeDDS(img, bcn.FormatDXT5)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tex.Write(&buf); err != nil {
		return err
	}

	return os.WriteFile(dst, buf.Bytes(), 0644)
}
