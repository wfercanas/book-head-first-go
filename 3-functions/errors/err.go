package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("height cannot be negative")
	fmt.Println(err.Error())
	
	err = fmt.Errorf("%.2f is not a valid height", -5.0)
	fmt.Println(err.Error())
	fmt.Println(err)
}