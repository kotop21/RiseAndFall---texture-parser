package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kotop21/Raf-assets-utiling/internal/matte-builder"
)

func main() {
	modeDiff := flag.Bool("diffuse", false, "Process only diffuse and ao maps")
	modeNormal := flag.Bool("normal", false, "Process only normal and specular maps")
	flag.Parse()

	args := flag.Args()

	if !*modeDiff && !*modeNormal {
		if len(args) < 4 {
			fmt.Println("Usage full: matte-builder <diffuse> <ao> <normal> <specular>")
			os.Exit(1)
		}

		imgDiff, err := mattebuilder.OpenImage(args[0])
		if err != nil {
			fmt.Printf("Error loading diffuse: %v\n", err)
			os.Exit(1)
		}
		imgAO, err := mattebuilder.OpenImage(args[1])
		if err != nil {
			fmt.Printf("Error loading ao: %v\n", err)
			os.Exit(1)
		}
		imgNormal, err := mattebuilder.OpenImage(args[2])
		if err != nil {
			fmt.Printf("Error loading normal: %v\n", err)
			os.Exit(1)
		}
		imgSpec, err := mattebuilder.OpenImage(args[3])
		if err != nil {
			fmt.Printf("Error loading specular: %v\n", err)
			os.Exit(1)
		}

		outDiff := mattebuilder.BakeAO(imgDiff, imgAO, 0.85)
		outNormal := mattebuilder.PackNormalSpec(imgNormal, imgSpec)

		diffOutName := mattebuilder.GetOutPath(args[0], "")
		normalOutName := mattebuilder.GetOutPath(args[2], "_bump")

		if err := mattebuilder.SavePNG(diffOutName, outDiff); err != nil {
			fmt.Printf("Error saving diffuse: %v\n", err)
			os.Exit(1)
		}
		if err := mattebuilder.SavePNG(normalOutName, outNormal); err != nil {
			fmt.Printf("Error saving normal: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if *modeDiff {
		if len(args) < 2 {
			fmt.Println("Usage: matte-builder --diffuse <diffuse> <ao>")
			os.Exit(1)
		}
		imgDiff, err := mattebuilder.OpenImage(args[0])
		if err != nil {
			fmt.Printf("Error loading diffuse: %v\n", err)
			os.Exit(1)
		}
		imgAO, err := mattebuilder.OpenImage(args[1])
		if err != nil {
			fmt.Printf("Error loading ao: %v\n", err)
			os.Exit(1)
		}
		outDiff := mattebuilder.BakeAO(imgDiff, imgAO, 0.85)
		diffOutName := mattebuilder.GetOutPath(args[0], "")
		if err := mattebuilder.SavePNG(diffOutName, outDiff); err != nil {
			fmt.Printf("Error saving diffuse: %v\n", err)
			os.Exit(1)
		}
	}

	if *modeNormal {
		if len(args) < 2 {
			fmt.Println("Usage: matte-builder --normal <normal> <specular>")
			os.Exit(1)
		}
		imgNormal, err := mattebuilder.OpenImage(args[0])
		if err != nil {
			fmt.Printf("Error loading normal: %v\n", err)
			os.Exit(1)
		}
		imgSpec, err := mattebuilder.OpenImage(args[1])
		if err != nil {
			fmt.Printf("Error loading specular: %v\n", err)
			os.Exit(1)
		}
		outNormal := mattebuilder.PackNormalSpec(imgNormal, imgSpec)
		normalOutName := mattebuilder.GetOutPath(args[0], "_bump")
		if err := mattebuilder.SavePNG(normalOutName, outNormal); err != nil {
			fmt.Printf("Error saving normal: %v\n", err)
			os.Exit(1)
		}
	}
}
