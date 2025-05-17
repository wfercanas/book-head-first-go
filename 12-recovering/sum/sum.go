package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(filename string) (*os.File, error) {
	fmt.Println("Opening:", filename)
	return os.Open(filename)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

func GetFloats(filename string) ([]float64, error) {
	file, err := OpenFile(filename)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file)

	var numbers []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}

func main() {
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	sum := 0.0
	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("Sum: %.2f\n", sum)
}
