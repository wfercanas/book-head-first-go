package main

import "fmt"

func main() {
	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 5)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", false)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", false)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", false)
	fmt.Printf("Percent sign: %%\n")
}