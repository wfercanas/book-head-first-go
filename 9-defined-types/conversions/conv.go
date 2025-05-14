package main

import "fmt"

type Liters float64
type Gallons float64

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 1 / 3.785)
}

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	carFuel += ToGallons(Liters(40))
	busFuel += ToLiters(Gallons(30))
	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)
}
