package mattefixer

import (
	"os"
	"path/filepath"
)

func FindTexconv() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(exePath)

	localPath := filepath.Join(dir, "texconv.exe")
	if _, err := os.Stat(localPath); err == nil {
		return localPath, nil
	}

	libPath := filepath.Join(dir, "lib", "texconv.exe")
	if _, err := os.Stat(libPath); err == nil {
		return libPath, nil
	}

	projectLibPath := filepath.Join(dir, "..", "lib", "texconv.exe")
	if _, err := os.Stat(projectLibPath); err == nil {
		return projectLibPath, nil
	}

	return "", os.ErrNotExist
}
