package main

import (
	"fmt"
	"log"
	"shared/keyboard"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade < 60 {
		status = "failing"
	} else {
		status = "passing"
	}

	fmt.Println("A grade of", grade, "is", status)
}