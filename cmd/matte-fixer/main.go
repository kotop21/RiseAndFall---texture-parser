package main

import (
	"fmt"
	"os"

	mattefixer "github.com/kotop21/Raf-texture-utils/internal/matte-fixer"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: matte-fixer.exe <file.png> [strength] [format]")
		fmt.Println("Drag and drop a .png file onto the executable to process with default settings.")
		os.Stdin.Read(make([]byte, 1))
		return
	}

	filePath, strength, format, err := mattefixer.ParseArgs(os.Args)
	if err != nil {
		fmt.Printf("[Error] %v\n", err)
		os.Stdin.Read(make([]byte, 1))
		os.Exit(1)
	}

	err = mattefixer.ProcessFile(filePath, strength, format)
	if err != nil {
		fmt.Printf("[Error] Failed to process file: %v\n", err)
		os.Stdin.Read(make([]byte, 1))
		os.Exit(1)
	}

	fmt.Println("[Success] File processed and saved as DDS successfully!")
}
