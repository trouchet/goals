package main

import "fmt"

// Function to add two integers and return the result
func add(x, y int) int {
	return x + y
}

// Function with multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

// Function with named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // Naked return, returns the named variables (x, y)
}

func main() {
	// Call the add function
	result := add(3, 7)
	fmt.Println("Result of addition:", result)

	// Call the swap function
	a, b := "Hello", "World"
	a, b = swap(a, b)
	fmt.Println("After swapping:", a, b)

	// Call the split function
	sum := 17
	x, y := split(sum)
	fmt.Printf("Splitting %d gives %d and %d\n", sum, x, y)
}
