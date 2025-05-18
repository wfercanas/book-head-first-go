package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func scanDirectory(dir string) error {
	content, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, item := range content {
		filePath := filepath.Join(dir, item.Name())
		if item.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}

	return nil
}

func main() {
	err := scanDirectory("dir")
	if err != nil {
		log.Fatal(err)
	}
}
