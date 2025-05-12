package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscriber) {
	fmt.Printf("Name: %s\n", s.name)
	fmt.Printf("Rate: %.2f\n", s.rate)
	fmt.Printf("Active: %t\n", s.active)
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 0.1
	s.active = true
	return &s
}

func applyDiscount(s *subscriber) {
	s.rate += .05
}

func main() {
	var person *subscriber = defaultSubscriber("Mike")
	applyDiscount(person)
	printInfo(person)
}
