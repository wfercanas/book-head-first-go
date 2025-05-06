package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func minmax(min float64, max float64, numbers ...float64) []float64 {
	var result []float64

	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}

	return result
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		log.Fatal(err)
	}

	return input
}

func getFloat() float64 {
	input := getInput()
	float, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	return float
}

func getFloats() []float64 {
	input := getInput()
	split := strings.Split(input, " ")
	var floats []float64
	for _, value := range split {
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal(err)
		}

		floats = append(floats, number)
	}

	return floats
}


func main() {
	fmt.Println("Welcome to minmax")
	fmt.Printf("Give me the min value: ")
	min := getFloat()
	
	fmt.Printf("Give me the max value: ")
	max := getFloat()
	
	fmt.Printf("Give me the list of values: ")
	values := getFloats()
	
	result := minmax(min, max, values...)
	fmt.Println("Result:", result)

}