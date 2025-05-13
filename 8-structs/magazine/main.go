package main

import (
	"fmt"
	"magazine/subscription"
)

func main() {
	var subscriber1 subscription.Subcriber
	subscriber1.Name = "Ryan"
	subscriber1.Rate = 0.1
	subscriber1.Active = true

	var subscriber2 subscription.Subcriber = subscription.Subcriber{Name: "Rob", Rate: 0.1, Active: false}
	subscriber3 := subscription.Subcriber{Name: "Kimi", Rate: 0.1, Active: true}

	employee := subscription.Employee{Name: "Joey", Salary: 60000}

	address := subscription.Address{Street: "ABC 123", City: "Miami", State: "FL", PostalCode: "123456"}
	subscriber1.HomeAddress = address
	subscriber2.HomeAddress = address
	subscriber2.HomeAddress.City = "Orlando"

	fmt.Println(subscriber1)
	fmt.Println(subscriber2)
	fmt.Println(subscriber3)
	fmt.Println(employee)
	fmt.Println(address)
}
