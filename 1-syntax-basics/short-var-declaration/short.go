package main

import "fmt"

func main() {
	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Damon"

	fmt.Println(customerName)
	fmt.Println("Has ordered", quantity, "sheets.")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}