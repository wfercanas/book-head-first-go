package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

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

	length := float64(len(numbers))
	var sum float64
	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("Average: %.2f\n", sum / length)
}