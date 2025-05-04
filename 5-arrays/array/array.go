package main

import (
	"fmt"
	"time"
)

func main() {
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	notes[3] = "fa"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[6]) // ""

	var primes [5]int
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])
	fmt.Println(primes[1])
	fmt.Println(primes[4]) // 0

	var codes [4]int = [4]int{11, 22, 33, 44}
	fmt.Println(codes[0]) // 11

	letters := [8]string{"a"}
	fmt.Println(letters[0]) // "a"
	fmt.Println(letters[1]) // ""

	books := [3]string{
		"The lord of the rings",
		"Perfect life",
	}
	fmt.Println(books)
	fmt.Printf("%#v\n", books)

	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	fmt.Println(dates[0])
	fmt.Println(dates[2]) // 0001-01-01 00:00:00 +0000 UTC
}