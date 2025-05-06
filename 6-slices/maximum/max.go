package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {	
		if number > max {
			max = number
		}
	}
	
	return max
}

func main() {
	args := os.Args[1:]
	var numbers []float64
	
	for _, arg := range args {
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, number)
	}

	max := maximum(numbers...)
	

	fmt.Printf("max: %.2f\n", max)
}