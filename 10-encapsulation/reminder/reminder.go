package main

import (
	"fmt"
	"log"
	"reminder/calendar"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(15)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date.Day())
	fmt.Println(date.Month())
	fmt.Println(date.Year())
}
