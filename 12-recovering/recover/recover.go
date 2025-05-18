package main

import "fmt"

func calmDown() {
	fmt.Println(recover())
}

func freakOut() {
	defer calmDown()
	panic("I'm freaking out!")
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
