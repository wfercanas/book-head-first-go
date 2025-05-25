package main

import (
	"html/template"
	"log"
	"os"
)

type Invoice struct {
	Name    string
	Paid    bool
	Charges []float64
	Total   float64
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	html, err := template.ParseFiles("bill.html")
	check(err)

	bill := Invoice{Name: "Mary Gibbs", Paid: false, Charges: []float64{12.03, 3.35, 0.19}, Total: 16.00}
	err = html.Execute(os.Stdout, bill)
	check(err)
}
