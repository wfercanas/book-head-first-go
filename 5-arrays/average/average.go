package main

import "fmt"

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0

	for _, value := range numbers {
		sum += value
	}

	count := float64(len(numbers))
	fmt.Printf("Average: %.2f\n", sum/count)
}