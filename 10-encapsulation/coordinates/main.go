package main

import (
	"coordinates/geo"
	"fmt"
	"log"
)

func main() {
	landmark := geo.Landmark{}

	err := landmark.SetName("The Googleplex")
	if err != nil {
		log.Fatal(err)
	}

	err = landmark.SetLatitude(37)
	if err != nil {
		log.Fatal(err)
	}

	err = landmark.SetLongitude(-110)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(landmark.Name())
	fmt.Println(landmark.Latitude())
	fmt.Println(landmark.Longitude())
}
