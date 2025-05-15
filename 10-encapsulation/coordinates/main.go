package main

import (
	"coordinates/geo"
	"fmt"
)

func main() {
	coordinates := geo.Coordinates{}
	coordinates.SetLatitude(37)
	coordinates.SetLongitude(110)
	fmt.Println(coordinates)
}
