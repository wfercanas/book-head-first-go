package main

import "fmt"

func main() {
	numbers := make(map[string]int)
	numbers["filled"] = 1
	fmt.Println(numbers["filled"], numbers["unfilled"]) // 1 0

	booleans := map[string]bool{}
	booleans["true"] = true
	fmt.Println(booleans["true"], booleans["false"]) // true false

	var words map[string]string
	fmt.Printf("%#v\n", words) // map[string]string(nil)
	words = make(map[string]string) // Without this line it doesn't work
	words["hello"] = "world"
	fmt.Println(words["hello"], words["none"]) // "world "
}