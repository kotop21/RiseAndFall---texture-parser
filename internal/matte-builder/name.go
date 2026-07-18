package mattebuilder

import (
	"path/filepath"
	"strings"
)

func GetOutPath(basePath, suffix string) string {
	dir := filepath.Dir(basePath)
	filename := filepath.Base(basePath)
	ext := filepath.Ext(filename)
	nameWithoutExt := strings.TrimSuffix(filename, ext)
	return filepath.Join(dir, nameWithoutExt+suffix+".png")
}
