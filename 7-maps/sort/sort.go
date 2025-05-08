package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]float64{"Alma": 50.2, "Rohit": 86.5, "Carl": 97.3}
	names := make([]string, 0)

	for name := range grades {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s has a grade of: %.2f\n", name, grades[name])
	}
}