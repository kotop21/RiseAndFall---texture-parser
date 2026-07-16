package convertor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	_ "github.com/woozymasta/bcn/dds"
)

type job struct {
	src  string
	dst  string
	name string
	ext  string
}

func Unpack(originalDir, editingDir string) error {
	entries, err := os.ReadDir(originalDir)
	if err != nil {
		return err
	}

	workers := getWorkerCount(entries)
	fmt.Printf("[Info] Using %d workers for unpacking\n", workers)

	jobs := make(chan job, len(entries))
	var wg sync.WaitGroup

	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				fmt.Printf("[Unpack] Started: %s\n", j.name)
				var err error
				switch j.ext {
				case ".tga":
					err = UnpackTGA(j.src, j.dst)
				case ".dds":
					err = UnpackDDS(j.src, j.dst)
				case ".sst":
					err = UnpackSST(j.src, j.dst)
				}

				if err != nil {
					fmt.Printf("[Error] %s: %v\n", j.name, err)
				} else {
					fmt.Printf("[Unpack] Done: %s\n", j.name)
				}
			}
		}()
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		ext := strings.ToLower(filepath.Ext(name))
		if ext != ".tga" && ext != ".dds" && ext != ".sst" {
			continue
		}

		src := filepath.Join(originalDir, name)
		dst := filepath.Join(editingDir, name+".png")
		jobs <- job{src: src, dst: dst, name: name, ext: ext}
	}
	close(jobs)
	wg.Wait()

	return nil
}
