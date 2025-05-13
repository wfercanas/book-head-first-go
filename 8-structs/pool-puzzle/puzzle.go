package main

import (
	"fmt"
	"puzzle/geo"
)

func main() {
	location := geo.Coordinates{Latitude: 30, Longitude: 100}
	fmt.Println(location.Latitude)
	fmt.Println(location.Longitude)
}
