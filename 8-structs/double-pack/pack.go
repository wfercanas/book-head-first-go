package main

import "fmt"

type part struct {
	description string
	count       int
}

func doublePack(item *part) {
	item.count *= 2
}

func main() {
	var item part
	item.description = "Fuses"
	item.count = 5
	doublePack(&item)
	fmt.Println(item.description)
	fmt.Println(item.count)
}
