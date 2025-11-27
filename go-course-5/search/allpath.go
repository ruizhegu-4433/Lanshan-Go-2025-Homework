package search

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func Allpath(root string) ([]string, error) {
	var paths []string
	var rootPath string
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}

	rootPath, err := filepath.Abs(root)
	if err != nil {
		return paths, err
	}

	var walkpath func(dir string)
	walkpath = func(dir string) {
		defer wg.Done()
		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println("Error accessing path:", dir, "error:", err, "skip")
			return
		}
		for _, entry := range entries {
			fullPath := filepath.Join(dir, entry.Name())
			lock.Lock()
			paths = append(paths, fullPath)
			lock.Unlock()
			if entry.IsDir() {
				wg.Add(1)
				go walkpath(fullPath)
			}
		}

	}

	wg.Add(1)
	walkpath(rootPath)
	wg.Wait()
	return paths, nil
}
