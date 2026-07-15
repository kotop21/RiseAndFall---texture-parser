package main

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/ftrvxmtrx/tga"
	"github.com/woozymasta/bcn"
	_ "github.com/woozymasta/bcn/dds"
)

var Version = "dev"

const (
	originalDir = "1_original"
	editingDir  = "2_editing"
	packedDir   = "3_packed"
)

func main() {
	if len(os.Args) == 1 {
		_ = os.MkdirAll(originalDir, 0755)
		_ = os.MkdirAll(editingDir, 0755)
		_ = os.MkdirAll(packedDir, 0755)

		fmt.Println("Command unpack/pack: `./converter.exe pack`")
		fmt.Println("Created:")
		fmt.Println(" -", originalDir)
		fmt.Println(" -", editingDir)
		fmt.Println(" -", packedDir)
		return
	}

	mode := strings.ToLower(os.Args[1])

	switch mode {
	case "unpack":
		_ = os.MkdirAll(originalDir, 0755)
		_ = os.MkdirAll(editingDir, 0755)
		_ = os.MkdirAll(packedDir, 0755)

		if err := unpack(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "pack":
		_ = os.MkdirAll(originalDir, 0755)
		_ = os.MkdirAll(editingDir, 0755)
		_ = os.MkdirAll(packedDir, 0755)

		if err := pack(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	default:
		fmt.Println("converter")
		fmt.Println("converter unpack")
		fmt.Println("converter pack")
	}
}

func unpack() error {
	entries, err := os.ReadDir(originalDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		ext := strings.ToLower(filepath.Ext(name))

		src := filepath.Join(originalDir, name)
		dst := filepath.Join(editingDir, name+".png")

		switch ext {
		case ".tga":
			if err := unpackTGA(src, dst); err != nil {
				fmt.Printf("%s: %v\n", name, err)
			}
		case ".dds":
			if err := unpackDDS(src, dst); err != nil {
				fmt.Printf("%s: %v\n", name, err)
			}
		}
	}

	return nil
}

func pack() error {
	entries, err := os.ReadDir(editingDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if !strings.HasSuffix(strings.ToLower(name), ".png") {
			continue
		}

		base := strings.TrimSuffix(name, ".png")
		ext := strings.ToLower(filepath.Ext(base))

		src := filepath.Join(editingDir, name)
		dst := filepath.Join(packedDir, base)

		switch ext {
		case ".tga":
			if err := packTGA(src, dst); err != nil {
				fmt.Printf("%s: %v\n", name, err)
			}
		case ".dds":
			if err := packDDS(src, dst); err != nil {
				fmt.Printf("%s: %v\n", name, err)
			}
		}
	}

	return nil
}

func unpackTGA(src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	img, err := tga.Decode(f)
	if err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	return png.Encode(out, img)
}

func packTGA(src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	return tga.Encode(out, img)
}

func unpackDDS(src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	_, img, err := bcn.DecodeDDS(f)
	if err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	return png.Encode(out, img)
}

func packDDS(src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return err
	}

	tex, err := bcn.EncodeDDS(img, bcn.FormatDXT5)
	if err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	return tex.Write(out)
}
