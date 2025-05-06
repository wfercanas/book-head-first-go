package main

import "fmt"

func print(strings ...string) {
	fmt.Println(strings)
}

func main() {
	print("a")
	print("a", "b", "c")
	print()
}