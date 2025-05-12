package main

import "fmt"

type person struct {
	name string
}

func main() {
	var user person
	user.name = "John"
	var pointer *person = &user
	fmt.Println(pointer)      // prints: &{John}
	fmt.Println(pointer.name) // prints: John
	// fmt.Println(*pointer.name)  // error -> pointer.name variable of type string. Invalid operation.
	fmt.Println(*&pointer.name)  // prints: John
	fmt.Println(pointer.name)    // prints: John
	fmt.Println((*pointer).name) // prints: John
}
