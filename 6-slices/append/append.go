package main

import "fmt"

func main() {
	// Basic append
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))
	slice = append(slice, "d", "e")
	fmt.Println(slice, len(slice))

	fmt.Println("-----------------------")

	// Without reassignment
	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	fmt.Println(s1, s2, s3, s4)

	s4[0] = "XX" // This changes s3 as well!!! They share the same underlying array.
	fmt.Println(s1, s2, s3, s4)
}