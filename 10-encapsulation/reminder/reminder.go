package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) setYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date) setMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) setDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}

func main() {
	date := Date{}
	err := date.setYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	err = date.setMonth(5)
	if err != nil {
		log.Fatal(err)
	}

	err = date.setDay(15)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
}
