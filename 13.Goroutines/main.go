package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func pause() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

func sendMsg(msg string, wg *sync.WaitGroup) {
	defer wg.Done() // Mark as done when the goroutine completes
	pause()
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3) // We have three goroutines to wait for

	// Anonymous function for "test1"
	go func(msg string) {
		defer wg.Done()
		pause()
		fmt.Println(msg)
	}("test1")

	// Corrected function calls
	go sendMsg("test2", &wg)
	go sendMsg("test3", &wg)

	wg.Wait() // Wait for all goroutines to complete
}
