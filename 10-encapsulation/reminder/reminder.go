package main

import (
	"fmt"
	"log"
	"reminder/calendar"
)

func main() {
	event := calendar.Event{}

	err := event.SetTitle("Mom's birthday")
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(15)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Title())
	fmt.Println(event.Day())
	fmt.Println(event.Month())
	fmt.Println(event.Year())

}
