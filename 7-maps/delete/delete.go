package main

import "fmt"

func main() {
	ranks := make(map[string]int)
	const literal = "bronze"
	ranks[literal] = 3
	
	rank, ok := ranks[literal]
	fmt.Printf("%s: %d -- ok: %t\n", literal, rank, ok)

	delete(ranks, literal)

	rank, ok = ranks[literal]
	fmt.Printf("%s: %d -- ok: %t\n", literal, rank, ok)
}