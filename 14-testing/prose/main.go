package main

import (
	"fmt"
	"prose/join"
)

func main() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", join.JoinWithCommas(phrases))
	phrases = append(phrases, "a prize bull")
	fmt.Println("A photo of", join.JoinWithCommas(phrases))
}
