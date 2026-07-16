package convertor

import (
	"os"
	"runtime"
)

func getWorkerCount(entries []os.DirEntry) int {
	var maxSize int64
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		info, err := entry.Info()
		if err == nil && info.Size() > maxSize {
			maxSize = info.Size()
		}
	}

	cpus := runtime.NumCPU()

	if maxSize > 50*1024*1024 {
		if cpus > 4 {
			return 2
		}
		return 1
	}
	if maxSize > 15*1024*1024 {
		w := cpus / 2
		if w < 1 {
			return 1
		}
		return w
	}
	return cpus
}
