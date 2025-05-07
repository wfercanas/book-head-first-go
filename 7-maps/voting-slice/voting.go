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

	var names []string
	var votes []int

	for _, line := range lines {
		match := false
		for i, name := range names {
			if line == name {
				votes[i]++
				match = true
				break
			}
		}

		if !match {
			names = append(names, line)
			votes = append(votes, 1)
		}
	}
	
	for i, name := range names {
		fmt.Printf("%v: %d\n", name, votes[i])
	}
}