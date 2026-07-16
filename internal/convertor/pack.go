package convertor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func Pack(editingDir, packedDir string) error {
	entries, err := os.ReadDir(editingDir)
	if err != nil {
		return err
	}

	workers := getWorkerCount(entries)
	fmt.Printf("[Info] Using %d workers for packing\n", workers)

	jobs := make(chan job, len(entries))
	var wg sync.WaitGroup

	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				fmt.Printf("[Pack] Started: %s\n", j.name)
				var err error
				switch j.ext {
				case ".tga":
					err = PackTGA(j.src, j.dst)
				case ".dds":
					err = PackDDS(j.src, j.dst)
				case ".sst":
					err = PackSST(j.src, j.dst)
				}

				if err != nil {
					fmt.Printf("[Error] %s: %v\n", j.name, err)
				} else {
					fmt.Printf("[Pack] Done: %s\n", j.name)
				}
			}
		}()
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
		if ext != ".tga" && ext != ".dds" && ext != ".sst" {
			continue
		}

		src := filepath.Join(editingDir, name)
		dst := filepath.Join(packedDir, base)
		jobs <- job{src: src, dst: dst, name: name, ext: ext}
	}
	close(jobs)
	wg.Wait()

	return nil
}
