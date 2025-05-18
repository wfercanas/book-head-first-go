package main

import (
	"fmt"
	"math/rand"
)

func prizeAward() {
	doorNumber := rand.Intn(3) + 1
	if doorNumber == 1 {
		fmt.Println("Won a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("Won a car!")
	} else if doorNumber == 3 {
		fmt.Println("Won a phone")
	} else {
		panic("This should not happen!")
	}
}

func main() {
	prizeAward()
}
