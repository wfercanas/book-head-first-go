package main

import (
	"coordinates/geo"
	"fmt"
	"log"
)

func main() {
	coordinates := geo.Coordinates{}
	err := coordinates.SetLatitude(37)
	if err != nil {
		log.Fatal(err)
	}

	err = coordinates.SetLongitude(-110)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())
}
