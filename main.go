package main

import (
	"errors"
	"fmt"
)

var ErrParameterGreaterThanFive = errors.New("parameter's value must be greater than 5")

func sum(a, b int) (*int, error) {
	if a < 6 || b < 6 {
		return nil, ErrParameterGreaterThanFive
	}
	sum := a + b

	return &sum, nil
}

func main() {
	fmt.Println("Welcome")
}
