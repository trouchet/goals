package main

import (
	"errors"
	"fmt"
)

// Function that returns an error
func divide(x, y float64) (float64, error) {
	if y == 0 {
		// Return an error for division by zero
		return 0, errors.New("division by zero")
	}
	result := x / y
	return result, nil
}

func main() {
	// Example of error handling
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of division:", result)
	}

	// Handling multiple return values, including errors
	result, err = divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of division:", result)
	}
}
