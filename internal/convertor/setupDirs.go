package convertor

import "os"

func SetupDirs(originalDir string, editingDir string, packedDir string) {
	_ = os.MkdirAll(originalDir, 0755)
	_ = os.MkdirAll(editingDir, 0755)
	_ = os.MkdirAll(packedDir, 0755)
}
