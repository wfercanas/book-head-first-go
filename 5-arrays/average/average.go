package main

import (
	"average/datafile"
	"fmt"
	"log"
)

func main() {
	numbers, err := datafile.GetFloat("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, value := range numbers {
		sum += value
	}

	count := float64(len(numbers))
	fmt.Printf("Average: %.2f\n", sum/count)
}