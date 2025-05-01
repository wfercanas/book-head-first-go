package main

import "fmt"

func multiple() (string, string) {
	return "hello", "world"
}

func namedMultiple() (firstname string, lastname string) {
	return "fernando", "canas"
}

func main() {
	fmt.Println(multiple())
	fmt.Println(namedMultiple())
}