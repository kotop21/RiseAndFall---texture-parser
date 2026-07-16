package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	_ "github.com/woozymasta/bcn/dds"
)

var Version = "dev"

const (
	originalDir = "1_original"
	editingDir  = "2_editing"
	packedDir   = "3_packed"
)

type job struct {
	src  string
	dst  string
	name string
	ext  string
}

func main() {
	if len(os.Args) == 1 {
		setupDirs()
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
		setupDirs()
		if err := unpack(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "pack":
		setupDirs()
		if err := pack(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("converter")
		fmt.Println("converter unpack")
		fmt.Println("converter pack")
	}
}

func setupDirs() {
	_ = os.MkdirAll(originalDir, 0755)
	_ = os.MkdirAll(editingDir, 0755)
	_ = os.MkdirAll(packedDir, 0755)
}

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

func unpack() error {
	entries, err := os.ReadDir(originalDir)
	if err != nil {
		return err
	}

	workers := getWorkerCount(entries)
	fmt.Printf("[Info] Using %d workers for unpacking\n", workers)

	jobs := make(chan job, len(entries))
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				fmt.Printf("[Unpack] Started: %s\n", j.name)
				var err error
				switch j.ext {
				case ".tga":
					err = unpackTGA(j.src, j.dst)
				case ".dds":
					err = unpackDDS(j.src, j.dst)
				case ".sst":
					err = unpackSST(j.src, j.dst)
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

func pack() error {
	entries, err := os.ReadDir(editingDir)
	if err != nil {
		return err
	}

	workers := getWorkerCount(entries)
	fmt.Printf("[Info] Using %d workers for packing\n", workers)

	jobs := make(chan job, len(entries))
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				fmt.Printf("[Pack] Started: %s\n", j.name)
				var err error
				switch j.ext {
				case ".tga":
					err = packTGA(j.src, j.dst)
				case ".dds":
					err = packDDS(j.src, j.dst)
				case ".sst":
					err = packSST(j.src, j.dst)
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
