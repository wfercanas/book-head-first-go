package main

import (
	"fmt"
	"reflect"
)

func createIntPointer() *int {
	myInt := 1
	return &myInt
}

func createFloatPointer() *float64 {
	myFloat := 1.23456
	return &myFloat
}

func printValueFromBoolPointer(pointer *bool) {
	fmt.Println(*pointer)
}


func main() {
	amount := 6
	fmt.Println(amount) 										// 6
	fmt.Println(&amount)										// 0x1400000e108
	fmt.Println(reflect.TypeOf(amount)) 		// int
	fmt.Println(reflect.TypeOf(&amount))		// *int

	fmt.Println("-------------------")
	
	var myIntPointer *int
	myIntPointer = createIntPointer()
	fmt.Println(myIntPointer)											// 0x14.....
	fmt.Println(reflect.TypeOf(myIntPointer))			// *int
	fmt.Println(*myIntPointer)										// 1

	fmt.Println("-------------------")

	var myFloatPointer *float64
	fmt.Println(myFloatPointer)										// <nil>
	myFloatPointer = createFloatPointer()											
	fmt.Println(myFloatPointer)										// 0x14....
	fmt.Println(reflect.TypeOf(myFloatPointer))		// *float64
	fmt.Println(*myFloatPointer)									// 1.00001

	fmt.Println("-------------------")

	myBool := true
	fmt.Println(myBool)														// true
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)										// 0x14.....
	printValueFromBoolPointer(myBoolPointer)			// true
	*myBoolPointer = false
	printValueFromBoolPointer(&myBool)						// false
	printValueFromBoolPointer(myBoolPointer)			// false
}