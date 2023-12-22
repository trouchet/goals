package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func fetchURL(url string, wg *sync.WaitGroup, results chan string) {
	defer wg.Done()

	// Make an HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("Error fetching %s: %s", url, err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		results <- fmt.Sprintf("Error reading %s response: %s", url, err)
		return
	}

	// Send the result through the channel
	results <- fmt.Sprintf("URL: %s\nContent: %s", url, body)
}

func main() {
	// List of URLs to fetch
	urls := []string{
		"https://www.example.com",
		"https://www.example.org",
		"https://www.example.net",
	}

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a channel to collect results from goroutines
	results := make(chan string, len(urls))

	// Start a goroutine for each URL
	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg, results)
	}

	// Use another goroutine to close the results channel once all fetching goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Receive and print results
	for result := range results {
		fmt.Println(result)
	}
}
