package main

import "fmt"

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	fmt.Printf("%#v\n", myStruct)
	fmt.Println(myStruct)

	myStruct.number = 3.14
	myStruct.word = "hello"
	myStruct.toggle = true

	fmt.Println(myStruct)
}
