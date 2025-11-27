package main

import (
	"bufio"
	"fmt"
	"go-course-5/search"
	"os"
	"strings"
)

/*
自己看着ai一点点写的，有些地方看的半懂不懂，比如说文件处理那个包里面的方法
谢谢观看
*/
func main() {
	var root string
	var pattern string
	var paths []string
	var err error
	channelinput := make(chan string, 1)
	fmt.Println("Start")
	fmt.Println("inuput root path:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	root = strings.TrimSpace(scanner.Text())
	paths, err = search.Allpath(root)
	for err != nil {
		fmt.Println("Error:", err)
	}
	for {
		fmt.Println("input pattern to search:('///quit' to exit)")
		scanner.Scan()
		pattern = strings.TrimSpace(scanner.Text())
		channelinput <- pattern
		if pattern == "///quit" {
			fmt.Println("Exiting program.")
			fmt.Println("put any key to close")
			scanner.Scan()
			close(channelinput)
			return
		}
		go func() {
			pattern := <-channelinput

			dirs, filepaths, err := search.SearchPath(pattern, paths)
			if err != nil {
				fmt.Println("Error:", err)
			}

			info, err := os.Stat(pattern)
			if err == nil && info.IsDir() {
				fmt.Println("Found directories:")
				for _, dir := range dirs {
					fmt.Println(dir)
				}
			} else {
				fmt.Println("Found files:")
				for _, file := range filepaths {
					fmt.Println(file)
				}
			}
		}()
	}

}
