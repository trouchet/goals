// factorial.go
package main

// Factorial calculates the factorial of a non-negative integer
func Factorial(n int) int {
	if n < 0 {
		return -1 // Error: Factorial is undefined for negative numbers
	}
	if n == 0 {
		return 1 // Base case: 0! is defined as 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	// This file doesn't contain executable code, as it's intended for testing.
}
