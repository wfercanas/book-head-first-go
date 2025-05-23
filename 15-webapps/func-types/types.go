package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func divide(a, b int) float64 {
	return float64(a) / float64(b)
}

func main() {
	var greeter func()
	var operation func(int, int) float64
	greeter = sayHi
	// greeter = divide // fails!
	operation = divide
	// operation = sayHi // fails!
	greeter()
	fmt.Println(operation(1, 2))
}
