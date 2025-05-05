package main

import "fmt"

func main() {
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice)
	fmt.Println(boolSlice)

	var slice []int
	var array [3]int
	fmt.Printf("%#v\n", slice) // []int(nil)
	fmt.Printf("%#v\n", array) // [3]int{0, 0, 0}
	fmt.Println(len(slice)) // Works even when the slice value is nil
}