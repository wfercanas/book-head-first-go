package main

import "fmt"

func main() {
	grades := make(map[string]float64)
	grades["Alma"] = 74.2
	grades["Rohit"] = 85.6
	grades["Carl"] = 59.7

	for name, grade := range grades {
		fmt.Printf("%s has a grade of: %.2f\n", name, grade)
	}
}