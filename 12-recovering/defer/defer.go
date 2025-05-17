package main

import (
	"errors"
	"fmt"
	"log"
)

func Greeting() error {
	defer fmt.Println("Goobye!")
	fmt.Println("Hello!")
	return errors.New("I don't want to talk")
	fmt.Println("I'm glad you're ok")
	return nil
}

func main() {
	err := Greeting()
	if err != nil {
		log.Fatal(err)
	}
}
