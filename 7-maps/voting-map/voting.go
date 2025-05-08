package main

import (
	"fmt"
	"log"
	"voting/readfile"
)

func main() {
	lines, err := readfile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	votes := make(map[string]int)
	for _, line := range lines {
		votes[line]++
	}
	
	// for i, name := range votes {
	// 	fmt.Printf("%v: %d\n", name, votes[i])
	// }

	fmt.Println(votes)
}