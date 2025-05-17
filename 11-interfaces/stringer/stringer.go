package main

import "fmt"

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%.2f gal", g)
}

func main() {
	carFuel := Gallons(11.5)
	fmt.Println("The car has", carFuel)
}
