package main

import "fmt"

func main() {
	// Slices
	numbers := []int{1, 2, 3, 4, 5}

	// Append to a slice
	numbers = append(numbers, 6)

	fmt.Println("Slice:", numbers)

	// Slice a slice
	slice := numbers[1:4]
	fmt.Println("Sliced Slice:", slice)

	// Maps
	person := map[string]string{
		"name":    "Alice",
		"country": "Wonderland",
	}

	// Add a key-value pair to the map
	person["age"] = "28"

	fmt.Println("Map:", person)

	// Access a value in the map
	fmt.Println("Name:", person["name"])

	// Check if a key exists
	if _, exists := person["gender"]; exists {
		fmt.Println("Gender exists:", exists)
	} else {
		fmt.Println("Gender does not exist.")
	}
}
