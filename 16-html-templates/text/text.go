package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal()
	}
}

func main() {
	text := "Here's my template\n"
	tmp, err := template.New("test").Parse(text)
	check(err)
	err = tmp.Execute(os.Stdout, nil)
	check(err)
}
