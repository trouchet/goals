// main.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	}

	// Register the handler for the root ("/") path
	http.HandleFunc("/", handler)

	// Start the web server on port 8080
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
