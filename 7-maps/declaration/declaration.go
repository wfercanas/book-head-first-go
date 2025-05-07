package main

import "fmt"

func main() {
	var myMap map[string]int
	myMap = make(map[string]int)

	newMap := make(map[string]int)
	literal := map[string]int{"a": 1, "b": 2}
	empty := map[string]int{}

	fmt.Println(myMap, newMap, literal, empty)
}