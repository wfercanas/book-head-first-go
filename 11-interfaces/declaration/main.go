package main

import (
	"declaration/mypkg"
	"fmt"
)

func main() {
	var value mypkg.MyInterface
	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(120)
	fmt.Println(value.MethodWithReturnValue())
}
