package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"image"
	"image/png"
	"os"

	"github.com/ftrvxmtrx/tga"
	"github.com/woozymasta/bcn"
)

func unpackSST(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	if len(data) < 15 {
		return errors.New("invalid SST file: too small")
	}

	revision := data[0]
	unknown := data[14]
	body := data[15:]

	var img image.Image
	var decodeErr error

	bodyReader := bytes.NewReader(body)

	if revision == 0 && unknown == 0 {
		img, decodeErr = tga.Decode(bodyReader)
	} else if revision == 1 {
		_, img, decodeErr = bcn.DecodeDDS(bodyReader)
	} else {
		return errors.New("unsupported SST format version")
	}

	if decodeErr != nil {
		return decodeErr
	}

	var outBuf bytes.Buffer
	if err := png.Encode(&outBuf, img); err != nil {
		return err
	}

	return os.WriteFile(dst, outBuf.Bytes(), 0644)
}

func packSST(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}

	bounds := img.Bounds()

	var tgaBuf bytes.Buffer
	if err := tga.Encode(&tgaBuf, img); err != nil {
		return err
	}

	header := make([]byte, 15)
	header[0] = 0
	header[1] = 1
	header[2] = 1

	binary.LittleEndian.PutUint32(header[6:10], uint32(bounds.Dx()))
	binary.LittleEndian.PutUint32(header[10:14], uint32(bounds.Dy()))

	header[14] = 0

	var outBuf bytes.Buffer
	outBuf.Write(header)
	outBuf.Write(tgaBuf.Bytes())

	return os.WriteFile(dst, outBuf.Bytes(), 0644)
}
