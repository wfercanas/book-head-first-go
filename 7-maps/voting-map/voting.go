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
	
	for name, count := range votes {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}