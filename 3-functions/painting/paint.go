package main

import (
	"fmt"
	"log"
)

func paintNeeded(width, height float64) (float64, error) {
	if height < 0 {
		return 0, fmt.Errorf("%.2f is an invalid height", height)
	}
	if width < 0 {
		return 0, fmt.Errorf("%.2f is an invalid width", width)
	}

	area := width * height
	return area / 10, nil
}

func main() {
	var total, amount float64
	var err error

	amount, err = paintNeeded(4.2, 3.0)	
	if err != nil {
		log.Fatal(err)
	}
	total += amount
	fmt.Printf("%.2f liters needed\n", amount)

	amount, err = paintNeeded(5.2, 3.5)
	if err != nil {
		log.Fatal(err)
	}
	total += amount
	fmt.Printf("%.2f liters needed\n", amount)
	fmt.Printf("Total: %.2f liters\n", total)
}