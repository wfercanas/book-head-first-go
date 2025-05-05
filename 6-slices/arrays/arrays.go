package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice1 := array[0:3]
	slice2 := array[2:]
	array[3] = 99

	fmt.Println(array)
	fmt.Println(slice1)
	fmt.Println(slice2)
}