package main

import "fmt"

func main() { 
	fmt.Printf("%12s | %s\n", "Products", "Cost in cents")
	fmt.Println("-----------------------------")
	fmt.Printf("%12s | %8.2f\n", "Stamps", 500.0)
	fmt.Printf("%12s | %8d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %8d\n", "Tape", 99)
}