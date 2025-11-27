package search

import (
	"fmt"
	"os"
	"strings"
)

func SearchPath(pattern string, allpath []string) ([]string, []string, error) {
	var dirs []string
	var filepaths []string
	for _, path := range allpath {
		info, err := os.Stat(path)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		if !info.IsDir() {
			if strings.HasSuffix(path, pattern) {
				filepaths = append(filepaths, path)
			}
		} else {
			dirs = append(dirs, path)
		}
	}
	return dirs, filepaths, nil
}
