package main

import "fmt"

func negate(boolPointer *bool) {
	*boolPointer = !*boolPointer
}

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)

	lies := false
	negate(&lies)
	fmt.Println(lies)
}