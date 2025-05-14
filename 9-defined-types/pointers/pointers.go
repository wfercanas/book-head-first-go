package main

import "fmt"

type Number int

func (n *Number) Double() {
	*n *= 2
}

func (n Number) Triple() {
	n *= 3
}

func main() {
	number := Number(2)
	number.Double()
	fmt.Println(number)

	address := &number
	address.Triple()
	fmt.Println(number)
}
