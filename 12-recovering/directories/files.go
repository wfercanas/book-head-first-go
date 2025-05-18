package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}

	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	}
}

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
	defer reportPanic()
	scanDirectory("dir")
}
