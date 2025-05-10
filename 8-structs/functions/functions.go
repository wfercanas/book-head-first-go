package main

import "fmt"

type part struct {
	description string
	count       int
}

func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func minimumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100

	return p
}

func main() {
	bolts := minimumOrder("Hex Bolts")
	showInfo(bolts)
}
