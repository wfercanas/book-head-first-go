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

	fmt.Println(subscriber1)
	fmt.Println(subscriber2)
	fmt.Println(subscriber3)
}
