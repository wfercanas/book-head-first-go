package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func sayBye() {
	fmt.Println("Bye")
}

func twice(callback func()) {
	callback()
	callback()
}

func main() {
	var hello func() = sayHi
	goodbye := sayBye
	twice(hello)
	twice(goodbye)
}
