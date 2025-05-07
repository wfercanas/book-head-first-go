package main

import "fmt"

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade, ok := grades[name]
	
	if !ok {
		fmt.Printf("%s doesn't exist\n", name)
		return
	}

	if grade < 60 {
		fmt.Printf("%s is failing\n", name)
	}

}

func main() {
	status("Alma") // failing because its zero
	status("Chris") // failing because of zero value

	// check with "comma ok idiom"
	
}