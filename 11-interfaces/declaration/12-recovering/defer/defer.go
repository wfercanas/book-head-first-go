package main

import "fmt"

func Greeting() {
	defer fmt.Println("Goobye!")
	fmt.Println("Hello!")
	fmt.Println("I'm glad you're ok")
}

func main() {
	Greeting()
}
