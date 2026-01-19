package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}

	// channel to collect results
	results := make(chan string, len(urls))

	for _, url := range urls {
		go func(u string) {
			// Simulate HTTP request with random delay
			delay := time.Duration(rand.Intn(1000)) * time.Millisecond
			time.Sleep(delay)
			results <- fmt.Sprintf("Scraped %s in %v", u, delay)
		}(url)
	}

	// Collect all results
	for i := 0; i < len(urls); i++ {
		recievedUrl := <- results
		fmt.Println(recievedUrl)
	}
	close(results)



}
