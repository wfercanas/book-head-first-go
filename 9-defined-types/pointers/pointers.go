package main

import "fmt"

type Number int

func (n *Number) Double() {
	*n *= 2
}

func (n Number) Print() {
	fmt.Println(n)
}

func main() {
	number := Number(2)
	number.Double()
	fmt.Println(number)

	address := &number
	address.Print()
	fmt.Println(number)
}
