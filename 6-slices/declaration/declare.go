package main

import "fmt"

func main() {
	var numbers []int
	// numbers := []int{1, 2, 3}
	numbers = make([]int, 5)
	fmt.Println(len(numbers))
}