package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func main() {
	var hello func()
	hello = sayHi
	hello()
}
