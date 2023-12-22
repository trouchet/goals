package main

import "fmt"

func main() {
	// If statement
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// Switch statement
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday.")
	case "Tuesday":
		fmt.Println("It's Tuesday.")
	default:
		fmt.Println("It's another day.")
	}

	// For loop
	fmt.Println("Counting from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// While loop equivalent
	fmt.Println("Counting from 1 to 5 using a while loop:")
	counter := 1
	for counter <= 5 {
		fmt.Println(counter)
		counter++
	}
}
