package main

import "fmt"

type Liters float64
type Gallons float64

func main() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = 10.0
	busFuel = 240.0

	fmt.Printf("The car has %.2f liters\n", carFuel*3.785)
	fmt.Printf("The bus has %.2f gallons\n", busFuel*1/3.785)
}
