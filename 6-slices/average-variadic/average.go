package main

import "fmt"

func average(numbers ...float64) float64{
	var sum float64 = 0
	length := float64(len(numbers))
	for _, number := range numbers {
		sum += number
	}

	return sum / length
}

func main() {
	fmt.Println(average(50, 100))
}