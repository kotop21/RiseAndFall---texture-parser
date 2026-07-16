package mattefixer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ProcessFile(filePath string, strength int, format string) error {
	texconv, err := FindTexconv()
	if err != nil {
		return fmt.Errorf("texconv.exe not found: %w", err)
	}

	ext := strings.ToLower(filepath.Ext(filePath))
	if ext != ".png" {
		return fmt.Errorf("only .png source files are supported")
	}

	tmpDir := os.TempDir()
	baseName := strings.TrimSuffix(filepath.Base(filePath), ext)
	tmpPNG := filepath.Join(tmpDir, baseName+"_matte_tmp.png")

	inputData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	err = os.WriteFile(tmpPNG, inputData, 0644)
	if err != nil {
		return fmt.Errorf("failed to create temporary PNG: %w", err)
	}

	err = ModifyAlpha(tmpPNG, strength)
	if err != nil {
		os.Remove(tmpPNG)
		return fmt.Errorf("failed to modify alpha channel: %w", err)
	}

	outputDir := filepath.Dir(filePath)
	err = RunTexconv(texconv, []string{
		"-f", format,
		"-m", "0",
		"-y",
		"-o", outputDir,
		tmpPNG,
	})

	os.Remove(tmpPNG)
	if err != nil {
		return fmt.Errorf("failed to compile to DDS: %w", err)
	}

	originalDDSName := baseName
	if strings.HasSuffix(strings.ToLower(originalDDSName), ".dds") {
		originalDDSName = strings.TrimSuffix(originalDDSName, filepath.Ext(originalDDSName))
	}

	generatedDDS := filepath.Join(outputDir, baseName+"_matte_tmp.dds")
	finalDDS := filepath.Join(outputDir, originalDDSName+".dds")

	if _, err := os.Stat(generatedDDS); err == nil {
		os.Remove(finalDDS)
		err = os.Rename(generatedDDS, finalDDS)
		if err != nil {
			return fmt.Errorf("failed to rename output file: %w", err)
		}
	}

	return nil
}
