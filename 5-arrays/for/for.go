package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "sol", "la", "si"}
	for i := 0; i < 2; i++ {
		fmt.Println(i, notes[i])
	}

	primes := [4]int{2, 3, 5, 7}
	for i := 0; i < len(primes); i++ {
		fmt.Println(i, primes[i])
	}

	for index, note := range notes {
		fmt.Println(index, note)
	}
}