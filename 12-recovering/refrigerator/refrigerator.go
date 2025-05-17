package main

import (
	"fmt"
	"log"
)

func find(item string, list []string) bool {
	for _, li := range list {
		if item == li {
			return true
		}
	}
	return false
}

type Refrigerator []string

func (r Refrigerator) Open() {
	fmt.Println("Opening refrigerator")
}

func (r Refrigerator) Close() {
	fmt.Println("Closing refrigerator")
}

func (r Refrigerator) Find(food string) error {
	r.Open()
	defer r.Close()

	lookup := find(food, r)
	if lookup == true {
		fmt.Println("Found", food)
	} else {
		return fmt.Errorf("%s not found", food)
	}

	return nil
}

func main() {
	fridge := Refrigerator{"Pizza", "Milk", "Cheese"}
	meal := []string{"Milk", "Bananas"}
	for _, food := range meal {
		err := fridge.Find(food)
		if err != nil {
			log.Fatal(err)
		}
	}
}
