package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return make([]float64, 0), err
	}

	floats := make([]float64, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return make([]float64, 0), err
		}

		floats = append(floats, number)	
	}

	err = file.Close()
	if err != nil {
		return make([]float64, 0), err
	}

	if scanner.Err() != nil {
		return make([]float64, 0), err
	}

	return floats, nil
}