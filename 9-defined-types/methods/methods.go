package main

import "fmt"

type MyType string

func (m MyType) MethodWithParameters(i int, s string) {
	fmt.Println(m)
	fmt.Println(i)
	fmt.Println(s)
}

func main() {
	value := MyType("Hello")
	value.MethodWithParameters(5, "World")
}
