package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func scanDirectory(dir string) {
	content, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, item := range content {
		filePath := filepath.Join(dir, item.Name())
		if item.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func main() {
	scanDirectory("dir")
}
