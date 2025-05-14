package main

import "fmt"

type Liters float64
type Gallons float64
type Title string

func main() {
	fmt.Println(Liters(10) + Liters(5))
	fmt.Println(Gallons(15) - Gallons(2))
	fmt.Println(Liters(1) * Liters(5))
	fmt.Println(Liters(6) == Liters(6))
	fmt.Println(6 == Liters(6))
	fmt.Println(Gallons(8) <= Gallons(4))
	fmt.Println(Gallons(10) + 8)

	fmt.Println("--------------")

	fmt.Println("Alien" == Title("Alien"))
	fmt.Println(Title("Alien") == Title("Alien"))
	fmt.Println(Title("Alien") > Title("Alien"))
	fmt.Println(Title("Alien") + Title("s"))
	fmt.Println(Title("Alien") + "s")
}
