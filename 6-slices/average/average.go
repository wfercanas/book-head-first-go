package main

import (
	"average/datafile"
	"fmt"
	"log"
)

func main() {
	floats, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0.0
	length := float64(len(floats))
	for _, number := range floats {
		sum += number
	}

	fmt.Printf("Average: %.2f\n", sum / length)
}