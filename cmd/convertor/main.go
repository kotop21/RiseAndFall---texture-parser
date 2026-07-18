package main

import (
	"fmt"
	"os"
	"strings"
)

var Version = "dev"

const (
	originalDir = "1_original"
	editingDir  = "2_editing"
	packedDir   = "3_packed"
)

func main() {
	if len(os.Args) == 1 {
		convertor.SetupDirs(originalDir, editingDir, packedDir)
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
		convertor.SetupDirs(originalDir, editingDir, packedDir)
		if err := convertor.Unpack(originalDir, editingDir); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "pack":
		convertor.SetupDirs(originalDir, editingDir, packedDir)
		if err := convertor.Pack(editingDir, packedDir); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("converter")
		fmt.Println("converter unpack")
		fmt.Println("converter pack")
	}
}
