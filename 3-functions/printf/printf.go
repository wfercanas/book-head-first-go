package main

import "fmt"

func main() {
	fmt.Println("About one third:", 1.0/3.0) // 0.33333333
	fmt.Printf("About one third: %f\n", 1.0/3.0) // 0.333333
	fmt.Printf("About one third: %.2f\n", 1.0/3.0) // 0.33

	fmt.Println("----------------------")

	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
	fmt.Printf("%%.1f: %.1f\n", 12.3456)
	fmt.Printf("%%.2f: %.2f\n", 12.3456)
}