package main

import (
	"fmt"
	"log"
	"shared/keyboard"
)

func main() {
	fmt.Print("Enter a Fahrenheit temperature: ")
	temp, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	celsius := (temp - 32) * 5 / 9
	fmt.Printf("%.2f Celsius degrees\n", celsius) 
}